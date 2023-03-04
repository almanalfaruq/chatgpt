package main

import (
	"fmt"
	"log"

	"github.com/almanalfaruq/chatgpt"
)

func main() {
	client, err := chatgpt.NewClient("MY_API_KEY", chatgpt.GPT35Turbo,
		"You're a travel assistant that know anything about travel and nothing else")
	if err != nil {
		log.Fatal(err)
	}

	reply, err := client.Chat("I want to enjoy the historical sites and cultural experience in Turkey")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Assitant:", reply)
	/*
		Example Output:
		Assitant: Great choice! Turkey has a rich history and culture that is definitely worth experiencing. Here are some must-see historical sites and cultural experiences you might want to consider:

		1. Hagia Sophia: This stunning cathedral-turned-mosque-turned-museum is a must-see while in Turkey. It was built in the 6th century and has served multiple purposes over the years. The intricate mosaics and colorful stained glass windows are a sight to behold.

		2. Topkapi Palace: This gigantic palace in Istanbul served as the home of the Ottoman sultans for centuries. You'll be able to see a wide variety of exhibits, including opulent costumes, royal jewels, and even the famous Topkapi dagger.

		3. Ephesus: This ancient Greek/Roman city is now a major archaeological site where you can explore ancient ruins, including the famous Library of Celsus and the Temple of Artemis.

		4. Whirling Dervish Ceremony: This performance is an authentic Sufi ritual that dates back hundreds of years. It's a mesmerizing display of spinning, chanting, and music that is sure to leave an impression.

		5. Turkish Bath: Also known as a hamam, this is a unique cultural experience that involves relaxing in a steam room, getting scrubbed down, and then massaged. It's a great way to unwind and experience a traditional Turkish ritual.

		6. Turkish Cuisine: Turkish food is known for its bold flavors and fresh ingredients. You'll definitely want to try some traditional dishes, such as kebabs, baklava, and Turkish delight.

		I hope these suggestions help you plan an amazing trip to Turkey!
	*/
}
