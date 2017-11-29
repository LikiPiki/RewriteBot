package main

var (
	admins = []string{
		"likipiki",
		"KirillQ",
		"websanya",
		"quantum_robin",
		"odintsov_design",
		"alexefremo",
	}
	keks = []string{
		"кек",
		"кпек",
		"КЕК",
		"Кек",
		"кееееееК",
		"лол кек чебурек",
		"ишак тебя метил",
		"КПЕК",
	}

	ornul = []string{
		"как мразь",
		"как тварь",
		"как пидор",
		"как шлеха",
		"как алеша попович",
		"как стерва",
		"как орало",
		"как орк",
		"как сосеска",
		"как барбареска",
	}

	replys = []string{
		"Вот ты мышь...",
		"Уйди лофтер!",
		"Ты надоел уже",
		"Шо, дизайнер шоли?",
		"Обкекался мразь",
		"Тоби пизда",
		"Вали пидор",
		"Соси писос",
		"Член не дорос кукарекать",
	}

	telki = []string{
		"Слишком много телок",
		"Попробуй позже",
		"Я отказываюсь выполнять твои команды",
	}

	telki_ne_daut = []string{
		"Иди работай!",
		"Телки только после работы",
		"Ну не в рабочее время же!",
		"Ты хочешь чтобы твои коллеги спалили тебя за телками? Я тоже нет",
		"Кончится работа - начнутся телки!",
		"Не в этот час",
		"Трудовые будни, не до телок",
	}

	errors = []string{
		"Телок дудосят",
		"Хватит фапать бро",
		"Что то пошло не так",
		"Алеша Попович занят с телками",
		"попробуй позже",
		"Азиатки кончились",
	}

	stikers = []string{
		"CAADAwADDwMAAt0NBwl-2xGCWqBcTAI", // pesiy kek
		"CAADAgADRwYAAkxb1gn9h6PpAyEkggI", // optimus fuck
		"CAADAgADngADk8vUCDsXkJ5Ka6VsAg",
		"CAADAgADIwADTFvWCfkn7nruQyHoAg",
		"CAADAgAD6gEAAsE8ngaA44zCtd3nBAI",
		"CAADAgADTQADZ0g6CQa56P8cfBcjAg", // spanch bob
	}
)

type JsonApiQuery struct {
	Preview string
	Id      int
}

type Querys []JsonApiQuery
