const fs = require('fs');
const text2png = require('text2png');

var text = `
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
`;

var buf = text2png(text, {
  font: '16px aahub',
  localFontPath: fs.realpathSync("./aahub.ttf"),
  textColor: "#333333",
  backgroundColor: "#ffffff",
  localFontName: 'aahub',
});
fs.writeFileSync("./out.png", buf);
