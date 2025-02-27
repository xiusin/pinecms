package helper

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"

	"xorm.io/xorm"

	"github.com/jinzhu/copier"
	"github.com/xiusin/pine"
	"github.com/xiusin/pine/di"
	"github.com/xiusin/pinecms/src/application/controllers"
	"github.com/xiusin/pinecms/src/application/models/tables"
)

type DirInfo struct {
	Label    string `json:"label"`
	FullPath string `json:"full_path"`
	IsDir    bool   `json:"is_dir"`
	Children any    `json:"children"`
}

const TimeFormat = time.DateTime

func init() {
	time.Local = time.FixedZone("CST", 8*3600)
	rand.Seed(time.Now().UnixNano())
}

func GetLocation() *time.Location {
	return time.Local
}

func AppPath() string {
	curPath, _ := os.Getwd()
	return curPath
}

// GetRootPath 获取项目根目录 (即 main.go的所在位置)
func GetRootPath(relPath ...string) string {
	pwd, _ := os.Getwd()
	if len(relPath) > 0 {
		pwd = filepath.Join(pwd, relPath[0])
	}
	return pwd
}

// GetCallerFuncName 获取当前执行函数名 只用于日志记录
func GetCallerFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	return runtime.FuncForPC(pc[0]).Name() + ":"
}

// Krand 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	for i := 0; i < size; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

// GetMd5 md5加密字符串
func GetMd5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}

// Ajax Ajax返回数据给前端
func Ajax(msg any, errcode int64, this *pine.Context) {
	if errcode == 0 {
		errcode = 1000
	}
	// 添加操作日志
	data := pine.H{"code": errcode}
	if errcode != 1000 {
		switch err := msg.(type) {
		case error:
			pine.Logger().Error("错误响应", err)
			data["message"] = err.Error()
		default:
			data["message"] = msg
		}
	} else {
		switch msg.(type) {
		case string:
			data["message"] = msg
		default:
			data["data"] = msg
		}
		data["data"] = msg
	}
	_ = this.Render().JSON(data)
}

// GetTimeStamp 获取时间戳
func GetTimeStamp() int {
	timestamp := time.Now().Unix()
	return int(timestamp)
}

// NowDate 当前时间 Y m d H:i:s
func NowDate(str string) string {
	return time.Now().Format(str)
}

// Password 生成密码
func Password(password, encrypt string) string {
	return GetMd5(GetMd5(password) + encrypt)
}

// IsFalse 检测字段是否为 空 0 nil
func IsFalse(args ...any) bool {
	for _, v := range args {
		switch v.(type) {
		case string:
			if v != "" {
				return false
			}
		case int, int64, int8, int32:
			if v != 0 {
				return false
			}
		case bool:
			if !v.(bool) {
				return false
			}
		default:
			return true
		}
	}
	return true
}

type EmailOpt struct {
	Title        string
	UrlOrMessage string
	Address      []string
}

func NewOrmLogFile(path string) *os.File {
	f, err := os.OpenFile(filepath.Join(path, "orm.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	PanicErr(err)
	return f
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetORM() *xorm.Engine {
	return pine.Make(controllers.ServiceXorm).(*xorm.Engine)
}

func ToInterfaces(values any) []any {
	v := reflect.ValueOf(values)
	if v.Kind() != reflect.Slice {
		return nil
	}
	var is []any
	for i := 0; i < v.Len(); i++ {
		is = append(is, v.Index(i).Interface())
	}
	return is
}

func DirTree(dir string) []DirInfo {
	fileInfos, _ := os.ReadDir(dir)
	var ms []DirInfo
	for _, f := range fileInfos {
		fullPath := filepath.Join(dir, f.Name())
		if f.IsDir() {
			s := DirInfo{
				Label:    f.Name(),
				IsDir:    true,
				FullPath: fullPath,
				Children: DirTree(fullPath),
			}
			ms = append(ms, s)
		} else {
			ext := strings.ToLower(filepath.Ext(f.Name()))
			if ext != ".css" && ext != ".js" && ext != ".jet" && ext != ".html" && ext != ".htm" && ext != ".sh" {
				continue
			}
			ms = append(ms, DirInfo{
				Label:    f.Name(),
				FullPath: fullPath,
				Children: "",
			})
		}
	}
	return ms
}

func UcFirst(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArr := []rune(str)
	if strArr[0] >= 97 && strArr[0] <= 122 {
		strArr[0] -= 32
	}
	return string(strArr)
}

func Bytes2String(b []byte) *string {
	return (*string)(unsafe.Pointer(&b))
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// 获取Url前缀
func GetUrlPrefix(catid int64) string {
	getUrlPrefix := di.MustGet(controllers.ServiceCatUrlPrefixFunc).(func(int64) string)
	return getUrlPrefix(catid)
}

// 处理文章列表信息数据. 补全一些cms生成
func HandleArtListInfo(list []map[string]string, titlelen int) {
	for i, art := range list {
		catid, _ := strconv.Atoi(art["catid"])
		prefix := GetUrlPrefix(int64(catid))
		if art["type"] != "2" {
			art["caturl"] = fmt.Sprintf("/%s/", prefix)
			art["typeurl"] = art["caturl"]
		}
		id, _ := strconv.Atoi(art["id"])
		art["arcurl"] = fmt.Sprintf("/%s/%d.html", prefix, id)
		art["arturl"] = art["arcurl"]
		art["click"] = art["visit_count"]
		art["fulltitle"] = art["title"]
		if titlelen > 0 {
			if utf8.RuneCountInString(art["title"]) > titlelen {
				titleRune := []rune(art["title"])
				art["title"] = string(titleRune[:titlelen])
			}
		}
		list[i] = art
	}
}

// PanicErr 抛出异常
func PanicErr(err error, msg ...string) {
	if err != nil {
		if len(msg) == 0 {
			panic(err)
		}
		panic(fmt.Sprintf("%s: %s", err, msg[0]))
	}
}

// Recover 捕捉执行协程的异常
func Recover() {
	if err := recover(); err != nil {
		pine.Logger().Error(fmt.Sprintf("%v", err), "stack", string(debug.Stack()))
	}
}

func ToTableTime(t string) tables.LocalTime {
	s, _ := time.ParseInLocation(time.DateTime, t, time.Local)
	return tables.LocalTime(s)
}

func ConvertToAnySlice[T any](s []T) []any {
	var slice = make([]any, 0, len(s))
	for _, v := range s {
		slice = append(slice, v)
	}
	return slice
}

// HasZero 常用类型零值判断
func HasZero(values ...any) bool {
	for _, value := range values {
		if IsZero(value) {
			return true
		}
	}
	return false
}

// IsZero 常用类型零值判断
func IsZero(t any) bool {
	if t == nil {
		return true
	}
	v := reflect.ValueOf(t)
	if !v.IsValid() {
		return true
	}
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return true
		}
		return IsZero(v.Elem().Interface())
	}
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Struct, reflect.Array:
		return v.Len() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	}
	return false
}

// ConvertTo 对结构体做copy转换
func ConvertTo[S any, T any](input *S, cbs ...func(*T, *S)) (*T, error) {
	data := new(T)
	if err := copier.Copy(data, input); err != nil {
		return nil, err
	}
	if cbs != nil {
		cbs[0](data, input)
	}
	return data, nil
}

// Or 常用类型default判断
func Or[T any](a T, b T) T {
	return CondOr(!IsZero(a), a, b)
}

// CondOr 常用类型default判断
func CondOr[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

// InArray 判断元素是否在切片中
func InArray[T comparable](item T, items []T) bool {
	for _, it := range items {
		if item == it {
			return true
		}
	}
	return false
}

// ArrayColumn 从items中获取元素值
func ArrayColumn[T any, I any](items []T, field string) []I {
	var values = make([]I, 0)
	var typeof reflect.Type
	var ok bool
	var fieldTypeof reflect.StructField

	for _, item := range items {
		if typeof == nil {
			typeof = reflect.TypeOf(item)
			if typeof.Kind() == reflect.Ptr {
				typeof = typeof.Elem()
			}
			if typeof.Kind() != reflect.Struct {
				panic(errors.New("ArrayColumn only support struct or struct pointer"))
			}
			if fieldTypeof, ok = typeof.FieldByName(field); !ok {
				for i := 0; i < typeof.NumField(); i++ {
					if tag, _ := typeof.Field(i).Tag.Lookup("json"); strings.HasPrefix(tag, field) {
						fieldTypeof = typeof.Field(i)
						break
					}
				}
			}
			if len(fieldTypeof.Name) == 0 {
				panic(fmt.Errorf("field [%s] not found", field))
			}
		}

		valueOf := reflect.ValueOf(item)
		if valueOf.Kind() == reflect.Ptr {
			if valueOf.IsNil() {
				continue
			}
			valueOf = valueOf.Elem()
		}

		if v, ok := valueOf.FieldByIndex(fieldTypeof.Index).Interface().(I); ok {
			v := v
			values = append(values, v)
		}
	}
	return values
}

// ArrayUnique 对切片去重
func ArrayUnique[T comparable](items []T) []T {
	exists := map[T]struct{}{}
	var filterItems []T
	for _, item := range items {
		if _, ok := exists[item]; !ok {
			exists[item] = struct{}{}
			filterItems = append(filterItems, item)
		}
	}
	return filterItems
}
