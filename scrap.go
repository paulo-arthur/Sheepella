package main

import (
	"log"
	"fmt"
	"context"
	//"strings"

	"github.com/chromedp/chromedp"
)

func scrap() {
	//Criando um contexto e um scraper
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	
	//Resultado bruto
	news := [5]string{}
	newsLinks := [5]string{}
	
	//Iniciando o brownser com tratamento de erro
	err := chromedp.Run(ctx, 
		chromedp.Navigate(`https://www1.folha.uol.com.br/internacional/en/scienceandhealth/`),
		chromedp.WaitVisible(`li.c-headline:nth-child(1) > div:nth-child(2) > div:nth-child(2) > a:nth-child(1) > h2:nth-child(1)`),

		chromedp.Text(`li.c-headline:nth-child(1) > div:nth-child(2) > div:nth-child(2) > a:nth-child(1) > h2:nth-child(1)`, &news[0]),
		chromedp.Text(`li.c-headline:nth-child(2) > div:nth-child(2) > div:nth-child(2) > a:nth-child(1) > h2:nth-child(1)`, &news[1]),
		chromedp.Text(`li.c-headline:nth-child(3) > div:nth-child(2) > div:nth-child(2) > a:nth-child(1) > h2:nth-child(1)`, &news[2]),
		chromedp.Text(`li.c-headline:nth-child(4) > div:nth-child(2) > div:nth-child(2) > a:nth-child(1) > h2:nth-child(1)`, &news[3]),
		chromedp.Text(`li.c-headline:nth-child(5) > div:nth-child(2) > div:nth-child(2) > a:nth-child(1) > h2:nth-child(1)`, &news[4]),

		chromedp.Evaluate(`document.evaluate("/html/body/main/div[1]/div[2]/div/div/div[1]/div/div/div/ol/li[4]/div[2]/div[2]/a", document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue.href;`, &newsLinks[0]),
		chromedp.Evaluate(`document.evaluate("/html/body/main/div[1]/div[2]/div/div/div[1]/div/div/div/ol/li[4]/div[2]/div[2]/a", document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue.href;`, &newsLinks[1]),
		chromedp.Evaluate(`document.evaluate("/html/body/main/div[1]/div[2]/div/div/div[1]/div/div/div/ol/li[4]/div[2]/div[2]/a", document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue.href;`, &newsLinks[2]),
		chromedp.Evaluate(`document.evaluate("/html/body/main/div[1]/div[2]/div/div/div[1]/div/div/div/ol/li[4]/div[2]/div[2]/a", document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue.href;`, &newsLinks[3]),
		chromedp.Evaluate(`document.evaluate("/html/body/main/div[1]/div[2]/div/div/div[1]/div/div/div/ol/li[4]/div[2]/div[2]/a", document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue.href;`, &newsLinks[4]),

		chromedp.Stop(),
	)
	
	//Tratando o erro
	if err != nil {
		log.Println(err)
		return
	}
	
	//Tratando o resultado, pois o mesmo vem no modelo:
	// YYYY/MM/DD HH-MM-SS||||resultado...
	// |#### INDESEJÁVEL ####|

	for key, new := range news {
		fmt.Printf("%v - %v\n", key + 1, new)
		fmt.Println(newsLinks[key])
		fmt.Println("")
	}

	//fmt.Printf(strings.Split(fmt.Sprintf("||||%q", news), "||||")[1])
	//fmt.Printf(strings.Split(fmt.Sprintf("||||%q", newsLinks), "||||")[1])
}