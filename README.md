Тестовое задание для компании Effective_Mobile

Условие 
Реализовать сервис, который будет получать по апи ФИО, из открытых апи обогащать
ответ наиболее вероятными возрастом, полом и национальностью и сохранять данные в
БД. По запросу выдавать инфу о найденных людях. Необходимо реализовать следующее
1 Выставить rest методы
1 Для получения данных с различными фильтрами и пагинацией
2 Для удаления по идентификатору
3 Для изменения сущности
4 Для добавления новых людей в формате
{
"name": "Dmitriy",
"surname": "Ushakov",
"patronymic": "Vasilevich" // необязательно
}
2 Корректное сообщение обогатить
1 Возрастом - https://api.agify.io/?name=Dmitriy
2 Полом - https://api.genderize.io/?name=Dmitriy
3 Национальностью - https://api.nationalize.io/?name=Dmitriy
3 Обогащенное сообщение положить в БД postgres (структура БД должна быть создана
путем миграций)
4 Покрыть код debug- и info-логами
5 Вынести конфигурационные данные в .env

Как запустить:
1) Сделать клон репозитория
  git clone --recurse-submodules https://github.com/M437A/test_task_Effective_Mobile

2) Запустить докер файл
   
  docker-compose build
  docker-compose up -d
или
  bash run.sh

3) Запустить само приложение
   go run cmd/main.go

Описание Эндпоинтов

Добавление нового пользователя
	post: http://localhost:8080

	request: {
		"name": "<NAME>", (required)
		"surname": "<Surname>", (required)
		"Patronymic": "<Patronymic>" (optional)
	}

	response: {
		"id": <id>
		"name": "<NAME>",
		"surname": "<Surname>",
		"Patronymic": "<Patronymic>",
		"age": "<age>",
		"gender": "<gender>",
		"nationality": "<nationality>"
	}

 Обновление данных о пользователи
  put: http://localhost:8080

	request: {
		"id": <id> (required)
		"name": "<NAME>", (required)
		"surname": "<Surname>", (required)
		"Patronymic": "<Patronymic>" (optional)
		"age": <age> (optional)
		"gender": "<gender>" (optional)
		"nationality": "<nationality>" (optional)
	}

	response: {
		"id": <id>
		"name": "<NAME>",
		"surname": "<Surname>",
		"Patronymic": "<Patronymic>",
		"age": "<age>",
		"gender": "<gender>",
		"nationality": "<nationality>"

 Постаричное получение с филтрами
	get: http://localhost:8080

	request: {
    "page": <PAGE>, (required)
    "user_filter": {
        "name": "<NAME>", (optional)
        "surname": "<SURNAME>", (optional)
        "patronymic": "<PATRONYMIC>", (optional)
        "gender": "<GENDER>", (optional)
        "nationality": "<NATIONALITY>" (optional)
  	  }
	}

	response: {
		list:
			"id": <id>
			"name": "<NAME>",
			"surname": "<Surname>",
			"Patronymic": "<Patronymic>",
			"age": "<age>",
			"gender": "<gender>",
			"nationality": "<nationality>"
	}

 Удаление по айди

	delete: http://localhost:8080/{userId}

	request: empty

	response: "User was deleted"

При обнаружение ошибки, просьба написать мне на почту 
al.matsnev.01@gmail.com


