// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	cptContextMarker = "description"
	cptMatchString   = `Многие думают, что Lorem Ipsum - взятый с потолка псевдо-латинский динамичный, развивающийся. набор слов, 
	но это не совсем так. Его корни уходят в один фрагмент классической латыни 45 года н.э., то есть более двух тысячелетий назад. 
	Ричард МакКлинток, профессор латыни из колледжа Hampden-Sydney, штат Вирджиния, взял одно из самых странных слов в Lorem Ipsum, 
	consectetur, и занялся его поисками в классической латинской литературе. В результате он нашёл неоспоримый первоисточник Lorem Ipsum 
	в разделах 1.10.32 и 1.10.33 книги de Finibus Bonorum et Malorum (О пределах добра и зла), написанной Цицероном в 45 году н.э. Этот 
	трактат по теории этики был очень популярен в эпоху Возрождения. Первая строка Lorem Ipsum, Lorem ipsum dolor sit amet.., происходит 
	от одной из строк в разделе 1.10.32 Классический текст Lorem Ipsum, используемый с XVI века, приведён ниже. 
	Также даны разделы 1.10.32 и 1.10.33 inibus Bonorum et Malorum Цицерона и их английский перевод, сделанный H. Rackham, 1914 год.
	Его корни уходят в один фрагмент классической латыни 45 года н.э., то есть более двух тысячелетий назад. 
	Ричард МакКлинток, профессор латыни из колледжа Hampden-Sydney, штат Вирджиния, взял одно из самых странных слов в Lorem Ipsum, 
	consectetur, и занялся его поисками в классической латинской литературе. В результате он нашёл неоспоримый первоисточник Lorem Ipsum 
	в разделах 1.10.32 и 1.10.33 книги de Finibus Bonorum et Malorum (О пределах добра и зла), написанной Цицероном в 45 году н.э. Этот 
	трактат по теории этики был очень популярен в эпоху Возрождения. Первая строка Lorem Ipsum, Lorem ipsum dolor sit amet.., происходит 
	от одной из строк в разделе 1.10.32 Классический текст Lorem Ipsum, используемый с XVI века, приведён ниже. 
	Также даны разделы 1.10.32 и 1.10.33 inibus Bonorum et Malorum Цицерона и их английский перевод, сделанный H. Rackham, 1914 год.
`
)

var (
	cptConcurrentProcessor *ConcurrentProcessor = NewConcurrentProcessor()
)

func TestConcurrentProcessorFindMatch(t *testing.T) {
	var matchTask = NewMatchTask()
	matchTask.AddSentence(cptContextMarker, cptMatchString)

	rules, err := cptConcurrentProcessor.FindMatch(matchTask)

	// TODO: load expected rules.
	var cptRulesExpected = Rules{}

	assert.NoError(t, err, "Expecting no errors.")
	assert.Equal(t, cptRulesExpected, rules, "Expecting rules are match.")
}

func BenchmarkConcurrentProcessorFindMatch(b *testing.B) {
	var matchTask = NewMatchTask()
	matchTask.AddSentence("description", cptMatchString)

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		cptConcurrentProcessor.FindMatch(matchTask)
	}
}

// go test ./index/rule -bench FindMatch -benchmem -cpu 1
// 1801229 ns/op    196356 B/op    16249 allocs/op

// go test ./index/rule -bench FindMatch -benchmem -cpu 8
//  423094 ns/op    205692 B/op    16314 allocs/op

// word purification before stemming (gain): ~ 100000 ns/op / 20%
