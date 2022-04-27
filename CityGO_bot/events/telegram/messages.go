package telegram

// const msgHelp = `I can save and keep you pages. Also I can offer you them to read.

// In order to save the page, just send me al link to it.

// In order to get a random page from your list, send me command /rnd.
// Caution! After that, this page will be removed from your list!`

const msgHelp = `Я бот City Go, отправь мне код который ты нашел и я проверю подходит ли он к испытаниям в которых ты учавствуешь!

Если не удаётся активировать код, проверь привязку Телеграм к аккаунту на сайте CityGO.kz. 
Для этого надо зайти в профиль и посмотреть свои контактные данные, если информация о телеграме отсутствует, то впиши туда свой логин`

const msgHello = "Hi there! 👾\n\n" + msgHelp

const (
	msgUnknownCommand           = "Unknown command 🤔"
	msgUserNotFound             = `Пользователь не найден. Пожалуйста, привяжи свой Телеграм к аккаунту на сайте CityGO.kz. Твой логин:`
	msgCodeNotFound             = "Данный код не подошел ни к одному испытанию в которых ты учавствуешь :("
	msgCodeActivatedSuccesfully = "Отлично! Ты активировал код и прошел квест"
)
