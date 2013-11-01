package main

import (
	"fmt"
	"bytes"
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"log"
)


var url string = "http://news.oneindia.in/rss/news-business-fb.xml"

type entertainment struct{

	Title []string `xml:"channel>item>title"`
	Description []string `xml:"channel>item>description"`

}



//getting entertainment news
func getNews(feed string) (n *entertainment,err error){

	res,err:=http.Get(feed) 
	
	defer res.Body.Close()
	if err!=nil{
		return nil,err
	}
	b,err:=ioutil.ReadAll(res.Body) //read the binaries of the object
	
	if err!=nil{
		return nil,err
	}

	gossip:=new(entertainment)
	

	err = xml.NewDecoder(bytes.NewBuffer(b)).Decode(&gossip) //Decode and put it in Gossip 

	return gossip,err
}


func main(){
	gossip,err := getNews(url)
	if err!=nil{
		log.Fatalf("Log: ", err)
		return 
	}

	for _,v:= range gossip.Title {
		fmt.Println("\n",v)
	}
	for _,v:=range gossip.Description {
		fmt.Println("\n",v)
	}
}
