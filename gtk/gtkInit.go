package gtk

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

var label *gtk.Label

func GtkInit() {
	//gtk.Init(&os.Args)
	//
	//win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	//win.SetTitle("gtk test")
	//win.SetUSize(800, 600)
	//
	//button := gtk.NewButton()
	//button.SetUSize(200, 50)
	//button.SetLabel("点击按钮")
	//
	//layout := gtk.NewFixed()
	//layout.Put(button, 50, 50)
	//
	////layout.Show()
	////button.Show()
	////win.Show()
	//
	//label = gtk.NewLabel("点击计数：0")
	//label.SetUSize(100, 50)
	//
	//layout.Put(label, 200, 200)
	//win.Add(layout)
	//
	//closeBtn := gtk.NewButton()
	//closeBtn.SetLabel("关闭")
	//closeBtn.SetUSize(50, 50)
	//closeBtn.Connect("clicked", func() {
	//	gtk.MainQuit()
	//})
	//
	//layout.Put(closeBtn, 750, 0)
	//
	//win.ShowAll()
	//
	//num := 0
	//
	//button.Connect("clicked", clickAction, &num)
	//

	//glade
	//glade := gtk.NewBuilder()
	//glade.AddFromFile("gladetest.glade")

	//win := gtk.
	//win := gtk.WindowFromObject(glade.GetObject("window1"))
	//btn := gtk.ButtonFromObject(glade.GetObject("close"))

	//gtk.Main()
}
func clickAction(sender *glib.CallbackContext) {
	args := sender.Data()
	msg, ok := args.(*int)
	if ok {
		*msg++
		fmt.Println(fmt.Sprintf("%d", *msg))
		label.SetLabel(fmt.Sprintf("点击计数：%d", *msg))
	}
}
