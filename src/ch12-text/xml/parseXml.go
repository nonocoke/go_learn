package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)


/*
解析XML到struct的时候遵循如下的规则:
  1.如果struct的一个字段是string或者[]byte类型且它的tag含有",innerxml"，Unmarshal将会将此字段所对应的元素内
    所有内嵌的原始xml累加到此字段上，如上面例子Description定义。最后的输出是
	  <server>
	  	  <serverName>Shanghai_VPN</serverName>
		  <serverIP>127.0.0.1</serverIP>
	  </server>
	  <server>
		  <serverName>Beijing_VPN</serverName>
		  <serverIP>127.0.0.2</serverIP>
	  </server>
  2.如果struct中有一个叫做XMLName，且类型为xml.Name字段，那么在解析的时候就会保存这个element的名字到该字段,如上面例子中的servers。
  3.如果某个struct字段的tag定义中含有XML结构中element的名称，那么解析的时候就会把相应的element值赋值给该字段，如上servername和serverip定义。
  4.如果某个struct字段的tag定义了中含有",attr"，那么解析的时候就会将该结构所对应的element的与字段同名的属性的值赋值给该字段，如上version定义。
  5.如果某个struct字段的tag定义 型如"a>b>c",则解析的时候，会将xml结构a下面的b下面的c元素的值赋值给该字段。
  6.如果某个struct字段的tag定义了"-",那么不会为该字段解析匹配任何xml数据。
  7.如果struct字段后面的tag定义了",any"，如果他的子元素在不满足其他的规则的时候就会匹配到这个字段。
  8.如果某个XML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有",comments"的字段上，
    这个字段的类型可能是[]byte或string,如果没有这样的字段存在，那么注释将会被抛弃。
 */
// 解析xml
// func Unmarshal(data []byte, v interface{}) error
func unpackXml() {

	type server struct {
		XMLName    xml.Name `xml:"server"`
		ServerName string   `xml:"serverName"`
		ServerIP   string   `xml:"serverIP"`
	}
	type Recurlyservers struct {
		XMLName     xml.Name `xml:"servers"`
		Version     string   `xml:"version,attr"`
		Svs         []server `xml:"server"`
		Description string   `xml:",innerxml"`
	}

	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

/*
生成的XML文件中的element元素名按照如下优先级从struct中获取:
  1.如果v是struct，XMLName的tag中定义的名称
  2.类型为xml.Name的名叫XMLName的字段的值
  3.通过struct中字段的tag来获取
  4.通过struct的字段名用来获取
  5.marshall的类型名称

如何设置struct 中字段的tag信息以控制最终xml文件的生成:
  - `XMLName不会被输出
  - tag中含有"-"的字段不会输出
  - tag中含有"name,attr"，会以name作为属性名，字段值作为值输出为这个XML元素的属性，如上version字段所描述
  - tag中含有",attr"，会以这个struct的字段名作为属性名输出为XML元素的属性，类似上一条，只是这个name默认是字段名了。
  - tag中含有",chardata"，输出为xml的 character data而非element。
  - tag中含有",innerxml"，将会被原样输出，而不会进行常规的编码过程
  - tag中含有",comment"，将被当作xml注释来输出，而不会进行常规的编码过程，字段值中不能含有"--"字符串
  - tag中含有"omitempty",如果该字段的值为空值那么该字段就不会被输出到XML，空值包括：false、0、nil指针或nil接口，
    任何长度为0的array, slice, map或者string
  - tag中含有"a>b>c"，那么就会循环输出三个元素a包含b，b包含c，例如如下代码就会输出
	  FirstName string   `xml:"name>first"`
	  LastName  string   `xml:"name>last"`

	  <name>
	  <first>Asta</first>
	  <last>Xie</last>
	  </name>
 */

// struct -> xml
// func Marshal(v interface{}) ([]byte, error)
// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
func packXml() {

	type server struct {
		ServerName string `xml:"serverName"`
		ServerIP   string `xml:"serverIP"`
	}

	type Servers struct {
		XMLName xml.Name `xml:"servers"`
		Version string   `xml:"version,attr"`
		Svs     []server `xml:"server"`
	}


	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, " ", "    ")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	//xml.MarshalIndent或者xml.Marshal输出的信息都是不带XML头的，为了生成正确的xml文件，我们使用了xml包预定义的Header变量
	os.Stdout.Write([]byte(xml.Header))

	os.Stdout.Write(output)
}

func main() {
	unpackXml()
	fmt.Println("- - - - - -")
	fmt.Println("- - - - - -")
	packXml()
}
