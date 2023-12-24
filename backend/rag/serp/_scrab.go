package serp

/*

func _scrab(url string, r io.Reader) (resp SpiderResult) {
	resp.Url = url
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		resp.Error = err
		return
	}
	doc.Find("title").Each(func(i int, selection *goquery.Selection) {
		resp.Title = selection.Text()
	})
	doc.Find("meta[name=description]").Each(func(i int, selection *goquery.Selection) {
		resp.Description, _ = selection.Attr("content")
	})
	doc.Find("body").Each(func(i int, selection *goquery.Selection) {
		defaultRemover(selection)
		if strings.Contains(url, "geeksforgeeks.org") {
			geeksforgeeks(selection)
		}
		if strings.Contains(url, "stackoverflow.com") {
			stackOveflow(selection)
		}
		resp.Content = selection.Text()
		resp.Content = strings.TrimSpace(resp.Content)
	})
	return
}

func defaultRemover(selection *goquery.Selection) {
	selection.Find("footer").Remove()
	selection.Find("header").Remove()
	selection.Find("noscript").Remove()
	selection.Find("script").Remove()
	selection.Find("[class*=footer]").Remove()
	selection.Find("[class*=sidebar]").Remove()
	selection.Find("[class*=search]").Remove()
	selection.Find("[class*=nav]").Remove()
	selection.Find("a[href*=twitter.com]").Remove()
	selection.Find("a[href*=facebook.com]").Remove()
	selection.Find("a[href*=linkedin.com]").Remove()
	selection.Find("[id*=footer]").Remove()
	selection.Find("[id*=sidebar]").Remove()
	selection.Find("nav").Remove()
	selection.Find("style").Remove()
}

func geeksforgeeks(selection *goquery.Selection) {
	defaultRemover(selection)
	selection.Find("[class*=footer]").Remove()
	selection.Find("[class*=header-main__container]").Remove()
	selection.Find("[class*=header-main__slider]").Remove()
	selection.Find("[class*=header-sidebar__wrapper]").Remove()
	selection.Find("[class*=gfg-footer]").Remove()
	selection.Find("[class*=article--recommended]").Remove()
	selection.Find("[class*=cookie-consent]").Remove()
}

func stackOveflow(selection *goquery.Selection) {
	defaultRemover(selection)
	//remove id
	selection.Find("[id*=left-sidebar]").Remove()
	selection.Find("[class*=js-dismissable-hero]").Remove()
	selection.Find("[id*=post-form]").Remove()

}
*/
