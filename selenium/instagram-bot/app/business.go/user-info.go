package business

import (
	"app/system"
	"app/utils"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

const (
	USERNAME       = "your_username"
	PASSWORD       = "your_password"
	TARGET_PROFILE = "https://www.instagram.com/darksidebooks/"
)

func CollectUserInfo() {
	// Seed é usado para que o random realmente funcione
	rand.Seed(time.Now().UnixNano())

	// Instancia o WebDriver
	wd := system.WebDriverInit()
	// TODO: Descomentar este defer
	// defer wd.Quit()

	log.Println("Acessando o endereço https://www.instagram.com/")
	err := wd.Get("https://www.instagram.com/")
	if err != nil {
		log.Println(err)
	}
	wd.SetImplicitWaitTimeout(5 * time.Second)
	time.Sleep(utils.GetRandomTimeDuration(5000, 6000) * time.Millisecond)

	// Preenche o nome do usuário no input
	log.Println("Preenchendo o nome de usuário")
	emailInput, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"loginForm\"]/div/div[1]/div/label/input") // TODO: Verificar se esse XPATH realmente é bom, por causa dessa div[1]
	if err != nil {
		log.Println(err)
	}
	if err := emailInput.SendKeys(USERNAME); err != nil {
		log.Println(err)
	}
	wd.SetImplicitWaitTimeout(5 * time.Second)
	time.Sleep(utils.GetRandomTimeDuration(2000, 3000) * time.Millisecond)

	// Preenche a senha no input
	log.Println("Preenchendo a senha")
	passwordInput, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"loginForm\"]/div/div[2]/div/label/input") // TODO: Encontrar outra maneira de escrever este XPATH
	if err != nil {
		log.Println(err)
	}
	if err := passwordInput.SendKeys(PASSWORD); err != nil {
		log.Println(err)
	}
	wd.SetImplicitWaitTimeout(5 * time.Second)
	time.Sleep(utils.GetRandomTimeDuration(4000, 6000) * time.Millisecond)

	// Faz o login
	log.Println("Fazendo o login")
	form, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"loginForm\"]/div/div[3]/button/div")
	if err != nil {
		log.Println(err)
	}
	if err := form.Click(); err != nil {
		log.Println(err)
	}
	wd.SetImplicitWaitTimeout(10 * time.Second)
	time.Sleep(utils.GetRandomTimeDuration(5000, 6000) * time.Millisecond)

	// Acessa o perfil alvo
	log.Println("Acessando o perfil", TARGET_PROFILE)
	err = wd.Get(TARGET_PROFILE)
	if err != nil {
		log.Println(err)
	}
	wd.SetImplicitWaitTimeout(5 * time.Second)
	time.Sleep(utils.GetRandomTimeDuration(5000, 6000) * time.Millisecond)

	// Rola a página 4 vezes
	for i := 0; i < 4; i++ {
		log.Println("Rolando a página até o final. Quantidade de vezes:", i+1)
		_, err = wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", []interface{}{})
		if err != nil {
			log.Println(err)
		}
		time.Sleep(utils.GetRandomTimeDuration(5000, 6000) * time.Millisecond)
	}

	// Pega as referências de cada post
	log.Println("Coletando os links de cada post")
	var posts []string
	links, err := wd.FindElements(selenium.ByTagName, "a")
	if err != nil {
		log.Println(err)
	}
	for _, link := range links {
		res, err := link.GetAttribute("href")
		if err != nil {
			log.Println(err)
		}

		if strings.Contains(res, "/p/") {
			log.Println("Encontrado o post:", res)
			posts = append(posts, res)
		}
	}

	for _, post := range posts {
		time.Sleep(utils.GetRandomTimeDuration(2000, 3000) * time.Millisecond)
		log.Println("Acessando o post:", post)
		err := wd.Get(post)
		if err != nil {
			log.Println(err)
		}
		// Acessa as pessoas que curtiram o post
		log.Println("Visualizando pessoas que curtiram esse post")
		followersLink, err := wd.FindElement(selenium.ByXPATH, "//a[contains(@href,\"/liked_by\")]")
		if err != nil {
			log.Println(err)
		}

		if err := followersLink.Click(); err != nil {
			log.Println(err)
		}
		// Esse sleep é para tentar abrir novamente o modal, pois, as vezes, ocorre
		// um erro que fecha o modal
		time.Sleep(utils.GetRandomTimeDuration(1000, 2000) * time.Millisecond)

		if err := followersLink.Click(); err != nil {
			log.Println(err)
		}
		wd.SetImplicitWaitTimeout(10 * time.Second)
		time.Sleep(utils.GetRandomTimeDuration(4000, 6000) * time.Millisecond)

		// Rola para baixo no modal de pessoas que curtiram o post
		likeModal, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[6]/div/div/div[2]/div")
		if err != nil {
			log.Println(err)
		}
		var likeModalInteface []interface{}
		likeModalInteface = append(likeModalInteface, likeModal)

		// Pega o número de pessoas que curtiu a publicação
		log.Println("Coletando o número de pessoas que curtiram a publicação")
		numLikes, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"react-root\"]/section/main/div/div[1]/article/div[3]/section[2]/div/div[2]/a/span")
		if err != nil {
			log.Println(err)
		}

		numLikesStr, err := numLikes.Text()
		if err != nil {
			log.Println(err)
		}

		numLikesInt, err := utils.GetFormatedNum(numLikesStr)
		if err != nil {
			log.Println(err)
		}
		log.Println(numLikesInt, "pessoas curtiram a publicação")

		var counter int
		if numLikesInt >= 1000 {
			counter = 1000 / 5
		}

		if numLikesInt < 1000 {
			counter = numLikesInt / 5
		}

		// TODO: Remover
		counter = 20

		// Rola a página para baixo
		var allUserNames []string
		for i := 0; i < counter; i++ {
			log.Println("Rolando a página para baixo. Contador:", i+1, "de", counter)
			_, err = wd.ExecuteScript("arguments[0].scrollTop = arguments[0].scrollTop + arguments[0].offsetHeight;", likeModalInteface)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(utils.GetRandomTimeDuration(2000, 3000) * time.Millisecond)
			//time.Sleep(utils.GetRandomTimeDuration(500, 1500) * time.Millisecond)

			// Pega os usernames de quem curtiu
			userNames, err := likeModal.FindElements(selenium.ByTagName, "a")
			if err != nil {
				log.Println(err)
			}
			log.Println("Coletando referência dos usuários")
			for _, userName := range userNames {
				userHref, err := userName.GetAttribute("href")
				if err != nil {
					log.Println(err)
				}
				allUserNames = append(allUserNames, userHref)
			}
		}
		userHrefs := utils.SetNoDuplicatedValue(allUserNames)
		for _, value := range userHrefs {
			fmt.Println(value)
		}

		//_, err = wd.ExecuteScript("document.evaluate('/html/body/div[6]/div/div/div[2]/div', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).scrollTop = 500", []interface{}{})
		//if err != nil {
		//	log.Println(err)
		//}

		// body > div.RnEpo.Yx5HN > div > div > div.Igw0E.IwRSH.eGOV_.vwCYk.i0EQd
		// /html/body/div[6]/div/div/div[2]

		//WebElement element = driver.findElement(By.xpath("//div[@id='document-content']/*[last()]"));

		// Pegar todos os nomes de usuário

		//os.Exit(1)
		time.Sleep(utils.GetRandomTimeDuration(9000, 11000) * time.Millisecond)
	}

	time.Sleep(utils.GetRandomTimeDuration(9000, 11000) * time.Millisecond)

	/*

		wellcome, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"welcome-panel\"]/div/h2")
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(wellcome.Text())
		}

		teste, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"loginform\"]/p[1]/label")
		if err != nil {
			log.Println(err)
		} else {
			teste2, err := teste.Text()
			if err != nil {
				log.Println(err)
			}
			fmt.Println(teste2)
		}

		commentsBtn, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"menu-comments\"]/a")
		if err != nil {
			log.Println(err)
		}

		if err := commentsBtn.Click(); err != nil {
			log.Println(err)
		}
		wd.SetImplicitWaitTimeout(10 * time.Second)
		time.Sleep(10 * time.Second)
	*/
}
