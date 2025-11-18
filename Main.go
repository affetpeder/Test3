package main

import (
    "log"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

// Go tarafında izin ve MediaStore çağrısı Java tarafında yapılacak
func main() {
    a := app.New()
    w := a.NewWindow("Media List")

    infoLabel := widget.NewLabel("İzin bekleniyor...")
    w.SetContent(container.NewVBox(infoLabel))
    w.Resize(fyne.NewSize(400, 200))

    // Java tarafı izin aldıktan sonra bu fonksiyon çağrılacak
    go func() {
        mediaFiles := GetMediaFiles() // JNI ile Java’dan çağrılacak
        if len(mediaFiles) == 0 {
            infoLabel.SetText("Fotoğraf bulunamadı veya izin reddedildi")
        } else {
            infoLabel.SetText("Fotoğraflar alındı: " + mediaFiles[0])
            for _, f := range mediaFiles {
                log.Println(f)
            }
        }
    }()

    w.ShowAndRun()
}
