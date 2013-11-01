package main

import (
	"fmt"
	"bytes"
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"log"
)


//var url string = "http://www.rediff.com/rss/inrss.xml"
var url string = "http://news.oneindia.in/rss/news-business-fb.xml"
type entertainment struct{

	Title string `xml:"channel>title"`
	Description []string `xml:"channel>item>description"`

}




//getting entertainment news
func getNews(feed string) (n *entertainment,err error){

	res,err:=http.Get(feed)
	
	defer res.Body.Close()
	if err!=nil{
		return nil,err
	}

	b,err:=ioutil.ReadAll(res.Body)
	
	if err!=nil{
		return nil,err
	}

	gossip:=new(entertainment)
	
	//err=xml.Unmarshal(b,&gossip)
	err = xml.NewDecoder(bytes.NewBuffer(b)).Decode(&gossip)
	fmt.Println("the answer : %v",gossip)
	fmt.Println("ERROR :")
	fmt.Println(err)

	return gossip,err
}


func main(){
	gossip,err := getNews(url)
	if err!=nil{
		log.Fatalf("Log: %v", err)
		return 
	}

	fmt.Println(gossip.Title)
	for _,v:=range gossip.Description {
		fmt.Println("\n",v)
	}

}
		
