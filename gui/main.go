package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gitlab-auto-merge/conf"

	"gitlab-auto-merge/gui/fonts"
	"gitlab-auto-merge/models"
	"gitlab-auto-merge/platform"
	"gitlab-auto-merge/service"
	"log"
	"strconv"
	"strings"
	"time"
)

var p *service.Service

func main() {

	p = service.NewService(platform.NewGitlab())

	myApp := app.New()
	myApp.Settings().SetTheme(&MyTheme{})
	myApp.SetIcon(fyne.NewStaticResource("icon", fonts.Logo))
	myWindow := myApp.NewWindow("自动提交工具")

	tabs := container.NewAppTabs(
		container.NewTabItem("首页", createTask()),
		container.NewTabItem("设置", choice()),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)
	myWindow.Resize(fyne.NewSize(500, 300))
	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func choice() fyne.CanvasObject {
	baseUrlEntry := widget.NewEntry()
	tokenEntry := widget.NewPasswordEntry()
	gconf := conf.GetConfig()
	baseUrlEntry.Text = gconf.Parameter.BasicUrl
	tokenEntry.Text = gconf.Parameter.Token

	form := widget.NewForm(
		&widget.FormItem{Text: "gitlab地址", Widget: baseUrlEntry, HintText: "http://gitlab.com"},
		&widget.FormItem{Text: "token", Widget: tokenEntry},
	)

	form.OnSubmit = func() {
		var c = conf.Config{
			Parameter: conf.ParameterS{
				BasicUrl: baseUrlEntry.Text,
				Token:    tokenEntry.Text,
			},
		}
		err := conf.UpdateConfig(c)
		if err != nil {
			log.Println(err)
		}
		//重新初始化
		p = service.NewService(platform.NewGitlab())
	}
	form.OnCancel = func() {
		baseUrlEntry.Text = ""
		tokenEntry.Text = ""
		form.Refresh()
	}

	return container.NewVBox(form)
}

func createTask2() fyne.CanvasObject {

	// 创建CheckGroup
	checkGroup := widget.NewCheckGroup([]string{"Option 1", "Option 2", "Option 3"}, func(selected []string) {
		fmt.Println("Selected options:", selected)
	})
	gInfo, err := p.GetGroups()
	if err != nil {
		log.Println(err)
	}
	groupMap := make(map[string]int)
	var groupName []string
	for _, v := range gInfo {
		groupMap[v.Name] = v.ID
		groupName = append(groupName, v.Name)
	}

	group := widget.NewSelect(groupName, func(value string) {
		log.Println("Select set to", value)
	})

	// 添加CheckGroup到窗口的内容容器
	return container.NewVBox(
		checkGroup,
		group,
	)
}

func createTask() fyne.CanvasObject {
	baseUrlEntry := widget.NewEntry()
	tokenEntry := widget.NewPasswordEntry()

	projectIDsEntry := widget.NewEntry()   //项目id
	groupIDsEntry := widget.NewEntry()     //分组id
	SourceBranchEntry := widget.NewEntry() //源分支
	TargetBranchEntry := widget.NewEntry() //目标分支
	TitleEntry := widget.NewEntry()        //标题
	ReviewerIDEntry := widget.NewEntry()   //审核人
	IntervalTimeEntry := widget.NewEntry() //间隔时间

	form := widget.NewForm(
		&widget.FormItem{Text: "项目id", Widget: projectIDsEntry},
		&widget.FormItem{Text: "分组id", Widget: groupIDsEntry},
		&widget.FormItem{Text: "源分支", Widget: SourceBranchEntry},
		&widget.FormItem{Text: "目标分支", Widget: TargetBranchEntry},
		&widget.FormItem{Text: "标题", Widget: TitleEntry},
		&widget.FormItem{Text: "审核人id", Widget: ReviewerIDEntry},
		&widget.FormItem{Text: "间隔时间(秒)", Widget: IntervalTimeEntry},
	)

	form.OnSubmit = func() {
		intervalTime, _ := strconv.Atoi(IntervalTimeEntry.Text)

		var req = models.TaskAutoMarge{
			ProjectIDs:   arrStringToInt(projectIDsEntry.Text),
			GroupIDs:     arrStringToInt(groupIDsEntry.Text),
			SourceBranch: SourceBranchEntry.Text,
			TargetBranch: TargetBranchEntry.Text,
			Title:        TitleEntry.Text,
			ReviewerID:   arrStringToInt(ReviewerIDEntry.Text),
			IntervalTime: intervalTime,
			CreatedTime:  time.Now(),
			Enable:       false,
			Cancel:       nil,
		}

		p.DelTask()

		err := p.CreateAutoMargeTask(&req)
		if err != nil {
			log.Println(err)
		}
		//重新初始化
		p = service.NewService(platform.NewGitlab())
	}
	form.OnCancel = func() {
		baseUrlEntry.Text = ""
		tokenEntry.Text = ""
		form.Refresh()
	}

	return container.NewVBox(form)
}
func arrStringToInt(str string) []int {

	strs := strings.Split(str, ",")
	var ints = make([]int, 0, len(strs))
	for _, s := range strs {
		k, _ := strconv.Atoi(s)
		ints = append(ints, k)
	}
	return ints
}
