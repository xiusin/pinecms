import term
import time

start := time.now()
term.clear()
publish_dir := "publish"
if is_dir("${publish_dir}") {
	rmdir_all("${publish_dir}")!
}
mkdir("${publish_dir}")!
mut exe := "pinecms"
println(term.bold(term.ok_message("开始构建执行文件")))
system('CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${exe}')
if !is_file(exe) {
	panic(exe + "不存在")
}
system("upx -9 ${exe}")

cp(exe, "${publish_dir}/${exe}")!
rm(exe)!

println(term.bold(term.ok_message("构建执行完成")))

println(term.bold(term.ok_message("开始打包前端页面")))
$if windows {
	system(".\\build.bat")
} $else {
	system("cd admin && yarn build")
}

mkdir("${publish_dir}/admin/")!

cp_all("admin/dist/", "${publish_dir}/admin/dist/", true) or {
	println(err)
}
cp_all("resources/", "${publish_dir}/resources/", true) or {
	println(err)
}

if is_dir ("build.dSYM") {
    rmdir_all("build.dSYM")!
}
system("zip -q -r publish.zip publish")
use_time := time.since(start)
println(term.bold(term.ok_message("构建完成, 目录: ${use_time.seconds()}")))
