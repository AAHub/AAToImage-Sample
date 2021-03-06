# coding: utf-8

import sys
from PIL import Image, ImageFont, ImageDraw

# reload(sys)
# sys.setdefaultencoding("utf-8")
# sys.getdefaultencoding()

aa = '''\
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
'''


lines = aa.split("\n")

font = ImageFont.truetype("../fonts/Saitamaar.ttf", 17)
w,h = max(font.getsize(line) for line in lines)

imag = Image.new("RGB", (w, h*len(lines)), "#ffffff")
draw = ImageDraw.Draw(imag)

for i,line in enumerate(lines):
    draw.text((0, i*h), line, font=font, fill="#000000")

# 表示
imag.show()
