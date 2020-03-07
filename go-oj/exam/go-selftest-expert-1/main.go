package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 图书借阅系统
func main() {

	r := bufio.NewReader(os.Stdin)
	records := []Record{}
	for {
		input, _ := r.ReadString('\n')
		if input == "" || input == "\n" {
			break
		}
		input = strings.Trim(input, "\n")
		inputs := strings.Split(input, ",")
		fmt.Println("inputs:", inputs)
		//fmt.Println("inputs len:", len(inputs))

		// 如果数据个数不符，继续走
		if len(inputs) != 3 {
			continue
		}

		r := &Record{}
		price, _ := strconv.Atoi(inputs[0])
		r.BookPrice = price
		plan, _ := strconv.Atoi(inputs[1])
		r.PlanDays = plan
		real, _ := strconv.Atoi(inputs[2])
		r.RealDays = real

		records = append(records, *r)
	}
	fmt.Println("所有records:", records)
	fmt.Println(calcBalance(records))
}

// 借书记录
type Record struct {
	BookPrice int
	PlanDays  int
	RealDays  int
}

// 每个用户借阅卡初始充值金额为300元。
const TOTAL int = 300

// 借书规则：
// 1、借阅图书的租金扣除规则：
//	A、价格大于等于100元的图书，累积借出天数小于等于15天时每本每天租金为5元；累积借出天数大于15天时，大于15天的时间每本每天租金为3元。
//	B、价格大于等于50元小于100元的图书，累积借出天数小于等于15天时每本每天租金为3元；累积借出天数大于15天时，大于15天的时间每本每天租金为2元。
//	C、价格小于50元的每天租金为1元。
// 2、余额少于当前所借书的价格时，不能借当前的书，但可以继续借阅其他更便宜的书。
// 3、需要考虑借书期间图书累积借出天数变化导致的租金变化（比如原价为120元的图书，预借天数为25天，实际借阅时间25天，那么前15天租金为5元，后10天为3元，借书时扣除租金为105元），扣除的租金最大不超过当前所借书的价格。
// 4、借书时需要说明预借天数，超期还书时，每超期一天，额外扣除1元。
func calcBalance(records []Record) (balance int) {
	balance = TOTAL // 初始约
	rent := 0       // 每条借书记录计算出的租金
	for i, r := range records {
		fmt.Printf("第%v轮:\n", i+1)
		fmt.Printf("  本轮record: %+v\n", r)
		// 2、余额少于当前所借书的价格时，不能借当前的书，但可以继续借阅其他更便宜的书。
		if balance < r.BookPrice {
			fmt.Printf("  本轮余额：%v 小于图书价格：%v，不能借阅当前书籍\n", balance, r.BookPrice)
			continue
		}
		rent = 0
		if r.BookPrice < 50 {
			//	C、价格小于50元的每天租金为1元。
			rent = r.RealDays * 1
		} else if r.BookPrice >= 50 && r.BookPrice < 100 {
			// B、价格大于等于50元小于100元的图书，累积借出天数小于等于15天时每本每天租金为3元；累积借出天数大于15天时，大于15天的时间每本每天租金为2元。
			if r.RealDays > 15 {
				rent = 15*3 + (r.RealDays-15)*2
			} else {
				rent = r.RealDays * 3
			}
		} else if r.BookPrice >= 100 {
			// A、价格大于等于100元的图书，累积借出天数小于等于15天时每本每天租金为5元；累积借出天数大于15天时，大于15天的时间每本每天租金为3元。
			if r.RealDays > 15 {
				rent = 15*5 + (r.RealDays-15)*3
			} else {
				rent = r.RealDays * 5
			}

		}

		// 4、借书时需要说明预借天数，超期还书时，每超期一天，额外扣除1元。
		if r.RealDays > r.PlanDays {
			rent += r.RealDays - r.PlanDays
		}

		// 扣除的租金最大不超过当前所借书的价格。
		if rent > r.BookPrice {
			rent = r.BookPrice
		}

		fmt.Println("  本轮rent:", rent)
		// 余额扣除租金
		balance -= rent
		fmt.Println("  本轮balance:", balance)
	}

	fmt.Print("最终余额: ")
	return balance
}
