package crawler_test

import (
	"testing"

	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/utils/crawler"
)

func TestCrawlData(t *testing.T) {
	url := "https://vnexpress.net/messi-suarez-cuu-inter-miami-thoat-thua-4719944.html"
	result := crawler.CreateCrawler(1).CrawlData(url, []string{})
	expectedResult := models.ExtractedData{
		Id:    "4719944",
		Title: "Messi, Suarez cứu Inter Miami thoát thua ",
		Paragraph: []string{
			"Ghi bàn: Shaffelburg 4, 46 - Messi 52, Suarez 90+5",
			"Phút 52, Messi nhận bóng ở trước cấm địa từ đường chuyền của Luis Suarez, mà không cầu thủ nào kịp ập vào áp sát. Thủ quân đội khách đẩy bóng hai nhịp, rồi cứa lòng về góc trái quá tầm với của thủ môn Joe Willis, rút tỷ số xuống 1-2. Bàn thắng này có thể coi là tuyệt phẩm với các cầu thủ, nhưng lại là sở trường của Messi.",
			"Tưởng như Miami phải nhận thất bại đầu tiên mùa 2024, Suarez lên tiếng ở phút bù hiệp hai. Từ quả tạt như đặt của Sergio Busquets từ bên phải, Suarez nhảy lên đánh đầu mà không bị ai kèm, đưa bóng về góc cao gỡ hòa 2-2. Bàn thắng của Messi và Suarez còn quan trọng bởi CONCACAF Champions Cup vẫn còn giữ luật ghi bàn sân khách. Thầy trò Tata Martino chỉ cần hòa 0-0 hay 1-1 ở lượt về tối 13/3 cũng đủ để vào tứ kết.",
			"Những cầu thủ ghi dấu giày vào hai bàn của Inter Miami đều là huyền thoại Barca, với Messi, Suarez và Busquets, trong ngày Jordi Alba chấn thương nghỉ chơi. Dù là tân binh, Miami vẫn được đánh giá cao hơn nhờ các ngôi sao này. Tuy nhiên, chủ nhà Nashville lại ghi liền hai bàn ở đầu mỗi hiệp để dẫn 2-0, khiến đội khách rơi vào thế rượt đuổi.",
			"Trận đấu này cũng tái hiện chung kết Leagues Cup mùa trước, và Messi cũng tỏa sáng với bàn gỡ, dù có thời điểm trong hiệp một anh phải nhờ bác sĩ chăm sóc đùi sau chân phải. Ngoài tình huống này, anh có lúc xô xát với cầu thủ chủ nhà Anibal Godoy ở phút 67, vì không đá bóng ra ngoài biên khi đối thủ nằm sân. Ngay lập tức, Suarez lao tới đẩy mạnh vào ngực khiến Godoy phải tránh xa Messi.",
			"Cũng có lúc Messi khiến khán giả đội khách thót tim khi anh nằm sân sau pha vào bóng của Lukas MacNaughton. Tình huống này MacNaughton phá bóng lên nhưng chân tiếp tục vung theo quán tính, đạp trúng bắp chân trái Messi. Trong lúc nằm sân, siêu sao 37 tuổi liên tục đập tay xuống đất, nhưng sau đó anh vẫn có thể chơi tiếp nhờ được bác sĩ chăm sóc. Lúc Messi nằm sân, nhiều khán giả giơ điện thoại lên ghi lại cảnh này.",
			"Miami tiếp tục khởi đầu mùa giải khá tốt, khi họ vẫn đang dẫn đầu MLS khu vực miền Đông, và có lợi thế ở Champions Cup. Nếu vô địch Champions Cup, Miami sẽ giành suất dự FIFA Club World Cup 2024, nơi Messi có thể gặp đội vô địch Champions League.",
			"CONCACAF Champions Cup được coi như Champions League của Bắc Trung Mỹ, là giải vô địch các CLB vùng này. Đây là lần đầu Inter Miami dự giải và được vào thẳng vòng 1/8, nhờ chức vô địch Leagues Cup 2023 - giải đấu Messi là Vua phá lưới và Cầu thủ hay nhất. Inter Miami là một trong tám CLB Mỹ tại Champions Cup, ngoài ra là các đội Mexico, Costa Rica và Suriname.",
			"Xuân Bình",
		},
		CrawledUrl: "https://vnexpress.net/messi-suarez-cuu-inter-miami-thoat-thua-4719944.html",
		RelatedUrl: []string{
			"https://vnexpress.net/klopp-neu-khong-phai-messi-ban-phai-phong-ngu-4718859.html",
			"https://vnexpress.net/haaland-messi-phai-giai-nghe-de-nguoi-khac-duoc-xem-la-hay-nhat-4718867.html",
			"https://vnexpress.net/leboeuf-bo-dao-nha-chi-gianh-euro-neu-loai-ronaldo-4718329.html",
			"https://vnexpress.net/neymar-hy-vong-choi-cung-messi-lan-nua-4717782.html",
			"https://vnexpress.net/messi-lap-cu-dup-bang-dau-va-nguc-4717742.html",
		},
		Img: []models.ImgStruct{
			{
				Src:         "https://i1-thethao.vnecdn.net/2024/03/08/messi-nashville-jpeg-170987171-3925-7679-1709871926.jpg?w=680&h=0&q=100&dpr=1&fit=crop&s=uAwQw4NuvExaA469mdcAhg",
				Description: "Messi (trái) mừng bàn vào lưới Nashville trên sân GEODIS ở thành phố Nashville, bang Tennessee, Mỹ, lượt đi vòng 1/8 CONCACAF Champions Cup tối 7/3/3024. Ảnh: USA Today",
			},
			{
				Src:         "https://i1-thethao.vnecdn.net/2024/03/08/messi-miami-jpeg-4150-1709871926.jpg?w=680&h=0&q=100&dpr=1&fit=crop&s=urAwprFzxuirTh_ynpCk5A",
				Description: "Messi rê bóng qua người trong trận đấu. Ảnh: USA Today",
			},
		},
	}

	if !equalStructs(result, expectedResult) {
		t.Errorf("unexpected word count: got %v, want %v", result, expectedResult)
	}
}

func equalStructs(p1, p2 models.ExtractedData) bool {
	// Compare each field of the structs
	if p1.Id != p2.Id {
		return false
	}
	if p1.Title != p2.Title {
		return false
	}
	if !slicesEqual(p1.RelatedUrl, p2.RelatedUrl) {
		return false
	}
	if !slicesEqual(p1.Paragraph, p2.Paragraph) {
		return false
	}
	if !compareImgStruct(p1.Img, p2.Img) {
		return false
	}
	// If all fields are equal, return true
	return true
}

func compareImgStruct(p1, p2 []models.ImgStruct) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i, v := range p1 {
		if v != p2[i] {
			return false
		}
	}
	return true
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
