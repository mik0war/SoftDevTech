package data

import (
	"errors"
	"online-shop-API/types"
	"online-shop-API/utils"
)

var Products = []types.Product{
	{
		ID:          "1",
		Name:        `16" Ноутбук Infinix InBook Y3H MAX YL613H серебристый`,
		Image:       "https://c.dns-shop.ru/thumb/st1/fit/500/500/34c9ab07ad55c60811ab07a6ea44d046/f711255f72ae017fb81a9aeb1a651fd7865e90c66444f01acbd7362ec6359e8f.jpg.webp",
		Description: "Infinix InBook Y3H MAX YL613H – это ноутбук, который сочетает в себе высокую производительность и стильный дизайн. Он идеально подходит для работы и развлечений, предлагая пользователям удобство и надежность. Этот ноутбук также отличается хорошей автономностью, что делает его отличным выбором для тех, кто часто находится в движении.",
		Category: types.Category{
			ID:   "2",
			Name: "Ноутбуки",
		},
	},

	{
		ID:          "2",
		Name:        `14.1" Ноутбук DEXP Aquilon серебристый`,
		Image:       "https://c.dns-shop.ru/thumb/st4/fit/500/500/7a003cd73cce7cad2347da0a93d7d281/00531b8fd5741f428c0666e46e44347488eed017c0dc5aae2aeae5809796e3fc.jpg.webp",
		Description: "Ноутбук DEXP Aquilon выполнен в серебристом корпусе диагональю 14.1\". IPS-экран с матовым покрытием отображает картинку в качестве Full HD с цветовым охватом NTSC 45%. Изображение будет комфортным для восприятия, а на экране не будет бликов. За обработку и вывод на экран графики отвечает видеокарта Intel UHD Graphics 600 с поддержкой декодирования Quick Sync. Шторка на веб-камере защищает конфиденциальность.\nDEXP Aquilon использует SSD-диск объемом 256 ГБ для хранения цифровой информации. Он ускоряет запуск приложений и загрузку файлов. DDR4-память объемом 8 ГБ в тандеме с процессором Intel Celeron N4020C отвечает за вычислительные операции, обеспечивает стабильную работу системы. Предустановленная Windows 11 Home удобна для решения повседневных задач. Она предлагает не только инструменты, интуитивно понятный интерфейс, но и игровые функции. DirectStorage ускоряет загрузку игр, делая систему подходящей для геймеров.",
		Category: types.Category{
			ID:   "1",
			Name: "Ноутбуки",
		},
	},
	{
		ID:          "3",
		Name:        `14.1" Ноутбук Chuwi HeroBook Pro серый`,
		Image:       "https://c.dns-shop.ru/thumb/st1/fit/500/500/71bd437bf61ba9597413381eaaa6aa8b/d1d3c26f6f53c01276c18eced8b9057aebc07c6cd746b469d168c11f1ed2e75e.jpg",
		Description: "HeroBook Pro оснащен 14,1-дюймовым IPS-экраном с более широкими углами обзора и разрешением 1920x1080. Кроме того, ночной режим и цветовые настройки помогут снизить напряжение глаз. 5,75-дюймовый тачпад поддерживает различные сенсорные жесты, обеспечивая простое и интуитивное взаимодействие с ноутбуком.",
		Category: types.Category{
			ID:   "1",
			Name: "Ноутбуки",
		},
	},
}

func GetProductsData(limit int, offset int, categoryId string) []types.Product {
	return utils.Filter(Products[offset:limit+offset], func(product types.Product) bool {
		return categoryId == "all" || product.Category.ID == categoryId
	})
}

func GetDataLength() int {
	return len(Products)
}

func CreateNewProduct(product types.Product) types.Product {
	Products = append(Products, product)

	return Products[len(Products)-1]
}

func UpdateProductData(id string, product types.Product) (int, error) {
	for index, product := range Products {
		if product.ID == id {
			Products[index] = product
			return index, nil
		}
	}
	Products = append(Products, product)

	return len(Products), errors.New("product not found")
}

func DeleteProductData(id string) (int, error) {

	for index, product := range Products {
		if product.ID == id {
			Products = append(Products[:index], Products[index+1:]...)
			return index, nil
		}
	}

	return 0, errors.New("product not found")
}
