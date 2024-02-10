package builder

import (
	"time"

	"github.com/thirteenths/test-bmstu23/internal/domain"
)

type EventBuilder struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (b *EventBuilder) Build() *domain.Event {
	return &domain.Event{
		ID:          b.ID,
		Name:        b.Name,
		Description: b.Description,
		Date:        b.Date,
	}
}

type EventMother struct{}

func (m EventMother) Obj0() *domain.Event {
	builder := EventBuilder{}
	return builder.Build()
}

func (m EventMother) Obj1() *domain.Event {
	builder := EventBuilder{
		ID:   1,
		Name: "Big Stand Up",
		Description: "Big Stand Up — шоу с самым большим процентом смеющихся людей. Здесь только опытные комики и шутки, проверенные не одной сотней избирательных зрителей." +
			" Приходите убедиться в пятницу, субботу и воскресенье, если вам больше 18 лет и вы свободны в пятницу, субботу и воскресенье.",
		Date: time.Time{},
	}
	return builder.Build()
}

func (m EventMother) Obj2() *domain.Event {
	builder := EventBuilder{
		ID:   2,
		Name: "Жёсткий стендап",
		Description: "Жёсткий стендап — это шоу, где комики могут шутить обо всем, о чем хотят, не боясь, что их сочтут сумасшедшими. " +
			"А зрители могут смеяться над всем, над чем хотят, не боясь, что это неуместно. Точно будут шутки про ХХX, не*******ю и к****c. " +
			"В шоу участвуют 4 комика и ведущий. Состав обновляется раз в месяц — успеете попасть на любимого стендапера. Приходите на шоу, чтобы посмеяться без стыда и, " +
			"возможно, расширить границы дозволенного. 18+",
		Date: time.Time{},
	}
	return builder.Build()
}

func (m EventMother) Obj3() *domain.Event {
	builder := EventBuilder{
		ID:   3,
		Name: "Женщины-комики",
		Description: "Женщины-комики — шоу, которое покажет силу женского юмора. В нём участвуют только девушки и только с лучшим своим материалом. " +
			"Три опытных комикессы, а также ведущая, расскажут качественные шутки. Как про мужчин, феминизм и психотерапию, так и про глобализацию, " +
			"энтропию и многое другое. Берите подругу, друга, да всех берите и приходите. Мы докажем вам, что женщины умеют шутить обо всём. 18+",
		Date: time.Time{},
	}
	return builder.Build()
}

func (m EventMother) Obj4() *domain.Event {
	builder := EventBuilder{
		ID:   4,
		Name: "Стендап Лайнап",
		Description: "айнап — это стендап-марафон без ведущего. Пять комиков рассказывают по 13 минут качественного материала. " +
			"В шоу вы можете увидеть опытных стендаперов с разным юмором. Также, в Лайнапе участвуют пока ещё не очень медийные комики, " +
			"но уже заработавшие себе авторитет в юмористическом комьюнити. Они покажут себя, познакомят с разными формами юмора и, возможно, " +
			"станут вашими любимчиками. Фишка шоу: если один комик не понравится, через 10 минут уже выступит следующий с другой подачей и отличающимся типом комедии. 18+",
		Date: time.Time{},
	}
	return builder.Build()
}
