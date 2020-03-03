const cheerio = require('cheerio')
const axios = require('axios')

const siteurl = "https://remoteok.io/"
let siteName ="";

const categories = new Set()
const tags = new Set()
const locations = new Set()
const positions = new Set()

const fetchData = async () => {
	const result = await axios.get(siteurl)
	return cheerio.load(result.data)
}

const getResults = async () => {
	const $ = await fetchData()
	siteName = $('.top > .action-post-job').text()

	$(".tags .tag").each((i, e) => {
		tags.add($(e).text())
	})

	$(".location").each((i, e) => {
		locations.add($(e).text())
	})

	$("div.nav p").each((i, e) => {
		categories.add($(e).text())
	})

	$('.company_and_position [itemprop="title"]').each((i, e) => {
		positions.add($(e).text())
	})

	// Convert everything to an array
	
	return {
		positions: [...positions].sort(),
		tags: [...tags].sort(),
		locations: [...locations].sort(),
		categories: [...categories].sort(),
		siteName,
	}
}

module.exports = getResults
