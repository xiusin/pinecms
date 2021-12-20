import term

term.clear()
println(term.bold(term.ok_message("1. 开始打包前端页面")))
if is_dir ("build") {
	rmdir_all("build") ?
}
mkdir("build")?
$if windows {
	system(".\\build.bat")
} $else {
	system("cd admin && yarn build")
}
cp_all("admin/dist/", "build/dist/", true) or {
	println(err)
}
mut exe := "pinecms"
$if windows {
	exe = "pinecms.exe"
}
println(term.bold(term.ok_message("2. 开始构建执行文件")))
system("go build -o ${exe}")
if !is_file(exe) {
	panic(exe + "不存在")
}
cp(exe, "build/${exe}")?
println(term.bold(term.ok_message("3. 复制配置文件")))
cp_all("resources/", "build/resources/", true) or {
	println(err)
}  