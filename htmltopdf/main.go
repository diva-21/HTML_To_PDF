package main

import (
    "fmt"
    "os"
    "sync"
    "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func htmlToPdfConvo(htmlFile string,wg *sync.WaitGroup ,pdfObj *wkhtmltopdf.PDFGenerator){
	defer wg.Done()
	// opening the file from dir
	file, err := os.Open(htmlFile)
	if err != nil {
		fmt.Println("Error opening HTML file:", err)
		return
	}
	defer file.Close()

	// Create PDF page from HTML file and add it to the PDF generator
	page := wkhtmltopdf.NewPageReader(file)
	pdfObj.AddPage(page)
}
func main() {
	// creating a pdf gen obj
	pdfObj, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        fmt.Println("Error creating PDF generator:", err)
        return;
    }
	pdfObj.Dpi.Set(300);
	pdfObj.NoCollate.Set(false)

	webPgs:=[]string{"./htmlFiles/pg1.html","./htmlFiles/pg2.html","./htmlFiles/pg3.html"}
	// creating a global var for WaitGroup
	var wg sync.WaitGroup
	for _,htmlFile:=range webPgs{
		wg.Add(1)
		go htmlToPdfConvo(htmlFile,&wg,pdfObj)
	}
	// Wait for all goroutines to finish
	wg.Wait()

	// Create PDF document in memory
	err = pdfObj.Create()
	if err != nil {
		fmt.Println("Error creating PDF document:", err)
		return
	}
// 	err = pdfObj.WriteFile("./htmlFiles/sample.pdf")
//   if err != nil {
//     fmt.Println("error in content write")
//   }

	file, errf := os.Create("output.pdf")
	if errf != nil {
		fmt.Println("Error creating output file:", errf)
		return
	}
	defer file.Close()
	// pdfString:=string(pdfObj.Bytes())
	// fmt.Println(pdfString)
	_, err = file.Write(pdfObj.Bytes())
	if err != nil {
		fmt.Println("Error writing PDF to file:", err)
		return
	}

	fmt.Println("PDF document generated successfully!")
}

