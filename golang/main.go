package main

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"image/png"
	"strings"
)

const (
	S           = 1024
	FONT_SIZE   = 16.0
	LINE_HEIGHT = 18.0
)

func main() {
	text := `
　 　 　 　 　 ／////////////////////////　------＼/////////ヽ
　　　　　　 ////:///////////////＞　　"´　　　　　　　 ∨////////∧
　　　　　 .///////////////＞ ´　　　　　　　　　　　 　 　 ∨////////∧
　　　　　/////////／/＞ ´　　　　　　　　　　　　　　　　 　 ∨////////∧
　　　　　|///////,〃／　　　　　　　　　　　　　　　　　　 　 　 i//////////ﾊ
　　　　　|////////i/　　　　　　　　　　　　_　　--――― ､--}//////////∧
　　　　　|////////|　　　　　　　　　　　 ´　　　　ﾄ.　　　　　ヽ ＼'/////////∧
　　　　　|////////|　　　　　 ..　 イ　　　　　　　 i　＼　　 ノ　 ',　＼////////∧
　　　　　|////////| 　 .　　 ´i　　 i　　 　 　 　 ｰ|-- 斗､.　　　 ',　　＼///////∧
　　　　　|////////| ／　　　 |　　ハ　　　 ',　 ∧!　,ｨf＝ミ､　　　',ー― ///////∧
　　　　　|////////|' 　　.、＿|..斗七ヽﾄ､ 　 ヽ| u 　 {::♥::} ゞ ＼　',i///////////,∧
　　　　　|////////|　　　　　 |/i/,.ィf=ミ ヽト､!|　　 ｀¨¨¨´/　　/＼!.ﾏニ ア//////,∧
　　　　　|////////ﾊ　　　 　/ .ｨ f:::♥ﾉ 　 | / ／／／／/　／　i　　|=イ/////////∧
　　　　　|////////,∧　 　 /{　ゞ ｀¨´ ／, j/　／／／ ./イ /　 .|　　i　 i///////////
　　　　　|/////////|.ﾍ　 /　乂　／／／／／　　　 　 u　/　　 |　　|　 i!///////////
　　　　　|/////////＞j／　　　＼ 　　　　　 　 　＿　　　j/|　　.| ∧!　　!///////////
　　　　　|//////／ ∠ｨ ﾍ 　　　　＞_ 　　 　 ,ｨv´　 _)　　/ |　　|/:::::∧ .|V//////////
　　　　　|//////ゝ---一 ﾍ　　ゝ----' u 　 ゝ-　´　 　 /i .|　　|:::::::::::∧!::∨/////////
　　　　　|//////////|　i　∧　　　　',≧=-　　 ＿＿　.イ　!ﾊ　　|:::::::::::::＿:::乂////////
　　　　　|//////////| ∧　∧　　　ト{ 　　　}　　　　i　　　 {ｰﾍ　!´￣　／ アニ≧= ---
　　　　　|//////////|/::::ﾍ.　 ',　　 |从　__ノi　　　　',　　　ﾉ　 ﾍ|　　／ ／ニニニニニニ
　　　　　|//////////|:::::::::::＼{ﾍ　 .|-＜　　 ゝ　　　 r ､ -､　　　　 /　 /ニニニニニニニ
　　　　　|/////////ノ==-＜　　 ',　|　　　　　　　　 r' 　 く　　　　　　./ニニニニニニﾆﾆ
　　　　 .ﾉ//＞≦ニニニニニ＼　 }/　　　　　　　 　 ｀ ´ ＼＼　　　　/ニニニニニニニニ
　　　 .//／ニニニニニニニニﾆ＼　　　　　　　　　　　　　 ｀´　　　　{ﾆﾆニニニニニニニ
　　　 .イニニニニニニニニニニニヽ　　　　　　　　　　　　　　　　　 /ニニニニニニニﾆﾆ
	`

	lines := strings.Split(text, "\n")

	if _, err := ConvertTextToImage(lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")
}

func ConvertTextToImage(lines []string) ([]byte, error) {
	// 対象アスキーアートの縦横を図る
	measure := gg.NewContext(S, S)
	if err := measure.LoadFontFace("../fonts/monapo.ttf", FONT_SIZE); err != nil {
		return nil, err
	}
	maxWidth := 0.0
	for _, line := range lines {
		w, _ := measure.MeasureString(line)
		if maxWidth <= w {
			maxWidth = w
		}
	}

	// 対象アスキーアートをpngに描画する
	width := int(maxWidth) + 10
	height := int(int(LINE_HEIGHT) * (len(lines) + 1))
	dc := gg.NewContext(width, height)
	if err := dc.LoadFontFace("./monapo.ttf", FONT_SIZE); err != nil {
		return nil, err
	}
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetHexColor("#333333")
	for idx, line := range lines {
		i := float64(idx + 1)
		dc.DrawString(line, 10, LINE_HEIGHT*i)
	}
	dc.Clip()
	dc.SavePNG("out.png")
	img := dc.Image()

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}
	ret := buf.Bytes()
	return ret, nil
}
