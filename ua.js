var data = []
document.querySelectorAll('ul > li > a').forEach((e) =>{
    data.push(e.text)
})
console.log(data.length)
var textFileAsBlob = new Blob([data.join("\r\n")], { type: "text/plain" }); //创建Blob对象，用于保存字符串为txt文件
var fileNameToSaveAs = "output.txt"; //指定输出文件的文件名
var downloadLink = document.createElement("a"); //创建一个链接元素，用于下载文件
downloadLink.download = fileNameToSaveAs; //链接元素的下载属性，指定要下载的文件名
downloadLink.innerHTML = "Download File";
downloadLink.href = window.webkitURL.createObjectURL(textFileAsBlob);
downloadLink.click(); //触发下载动作
