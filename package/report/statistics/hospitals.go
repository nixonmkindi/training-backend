package report

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/signintech/gopdf"
//	"eoffice/backend/package/log"
//	"eoffice/backend/package/util"
//	"eoffice/backend/webserver/models"
//	"time"
//)
//
//func HospitalReport(hospitals []*models.Hospital) (string, error) {
//
//	begin := time.Now()
//
//	contents := [][]string{}
//	tableTitle := []string{"S/N", "HOSPITAL NAME", "DISTRICT", "REGION", "HOSPITAL TYPE", "SERVICES"}
//	contents = append(contents, tableTitle)
//
//	//uploaded := 0
//	//notUploaded := 0
//	for k, hospital := range hospitals {
//		//className := []string{"CLASS :" + hospitals[k].Class.Name}
//		//if k == 0 || hospitals[k].Class.Name != hospitals[k-1].Class.Name {
//		//	contents = append(contents, className)
//		//}
//		//if status[k] == "Uploaded" {
//		//	uploaded++
//		//} else {
//		//	notUploaded++
//		//}
//		//staff := strings.ToUpper(module.Staff.LastName) + ", " + module.Staff.FirstName
//
//		row := []string{fmt.Sprintf("%d", k+1), hospital.HospitalName, hospital.District.Name, hospital.District.Region.Name, hospital.HospitalType, hospital.ConsultationServices.Name}
//		contents = append(contents, row)
//	}
//
//	mainTitle := "UNITED REPUBLIC OF TANZANIA"
//	reportTitle := "NATIONAL eoffice PLATFORM\n"
//	fileName := "Hospitals"
//
//	font := 12
//	//fileName := "module-result"
//	columnWidth := []float64{0.6, 1.2, 4.5, 1.2, 2.2, 1.2}
//	pdf := UploadStatusReportPages(mainTitle, reportTitle, contents, columnWidth, font, 1, true)
//
//	//add statistic page
//	//font = 12
//	pdf.AddPage()
//	currentPageNumber++
//	setPageNumb(pdf, currentPageNumber)
//	//setFontBold(pdf, font)
//
//	//add pie chart
//	chartWidth := pageWidth / 2.5
//	chartHeight := pageWidth / 2.5
//	//	uploadedPercentage := float64(uploaded) / float64(uploaded+notUploaded) * 100
//	//notUploadedPercentage := 100.0 - uploadedPercentage
//	//pie := chart.PieChart{
//	//	Width:  int(chartWidth),
//	//	Height: int(chartHeight),
//	//	Values: []chart.Value{
//	//		{Value: float64(uploaded), Label: fmt.Sprintf("UPLOADED(%.1f %%)", uploadedPercentage)},
//	//		{Value: float64(notUploaded), Label: fmt.Sprintf("NOT UPLOADED(%.1f %%)", notUploadedPercentage)},
//	//	},
//	//}
//	setFont(pdf, 11)
//	buffer := bytes.NewBuffer([]byte{})
//	err := pie.Render(chart.PNG, buffer)
//	if util.CheckError(err) {
//		log.Errorf("error rendering pie chart: %v", err)
//	}
//
//	addChart(pdf, pageWidth/2-chartWidth/2, headingMargin+padding, chartWidth, chartHeight, buffer.Bytes())
//
//	reportDir, err := config.ReportDir()
//	if util.CheckError(err) {
//		log.Errorf("error getting report directory %v", err)
//		return "", err
//	}
//	timeFileName := fmt.Sprintf("-%d", time.Now().Unix())
//	path := reportDir + fileName + timeFileName + ".pdf"
//	pdf.WritePdf(path)
//	pdf.Close()
//
//	end := time.Now()
//	fmt.Printf("PDF Report generated in %v\n", end.Sub(begin))
//	return path, nil
//}
//
//func UploadStatusReportPages(mainTitle, title string, data [][]string, columnWidth []float64, fontSize int, appendPages int, isLandcape bool) *gopdf.GoPdf {
//	d, err := json.Marshal(data)
//	if util.CheckError(err) {
//		log.Errorf("error getting json data %v", err)
//	}
//
//	qrs, doi, err := util.GetQRString(d)
//	if util.CheckError(err) {
//		log.Errorf("error getting qr string and doi %v", err)
//	}
//	//variable initalisation
//	//fmt.Printf("Is landscape: %v\n", isLandcape)
//	currentPageNumber = 1
//	totalPageNumber = appendPages
//	pdf := initPDF(isLandcape)
//	pdf.SetMargins(leftMargin, topMargin, rightMargin, bottomMargin)
//
//	//add fonts
//	addFonts(pdf)
//	pdf.AddPage()
//	pdf.SetX(rightMargin)
//	pdf.SetY(topMargin)
//	totalRows = float64(len(data)) //initalise the number of rows
//
//	//display header
//	//title = strings.ToUpper(title)
//	mainHeader(pdf, mainTitle, title, qrs, doi)
//
//	xp := leftMargin
//	xy := headingMargin
//
//	//normalise table width
//	totalWidth := 0.0
//	for _, w := range columnWidth {
//		totalWidth += w
//	}
//	for c, w := range columnWidth {
//		columnWidth[c] = w / totalWidth * availablePageWidth
//	}
//
//	tableHeading := data[0]
//	data = data[1:]
//
//	addTableleftAlign(pdf, xp, xy, 20, 5, columnWidth, tableHeading, data, fontSize)
//
//	return pdf
//}
