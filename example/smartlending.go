//go:build example

package main

import (
	"fmt"

	sdkpkg "github.com/whitebit-exchange/go-sdk/sdk"
)

func main() {
	// Create SDK with your API credentials
	s := sdkpkg.New("", "")

	fixedPlans, err := s.SmartLending.GetFixedPlans("USDT")
	if err != nil {
		fmt.Println("fixed plans error:", err)
	} else {
		fmt.Printf("Available plans: %d\n", len(fixedPlans))
		for _, plan := range fixedPlans {
			fmt.Printf("  ID: %s, Ticker: %s, Rate: %s%%, Duration: %d days\n",
				plan.ID, plan.Ticker, plan.Percent, plan.Duration)
			fmt.Printf("    Min: %s, Max: %s\n", plan.MinInvestment, plan.MaxInvestment)
		}
	}

	// Create a fixed investment
	fmt.Println("\n=== Create Fixed Investment ===")
	if len(fixedPlans) > 0 {
		investment, err := s.SmartLending.CreateFixedInvestment(fixedPlans[0].ID, "100")
		if err != nil {
			fmt.Println("create investment error:", err)
		} else {
			fmt.Printf("Investment ID: %s, Amount: %s, Status: %d\n",
				investment.ID, investment.Amount, investment.Status)
		}
	}

	// Get fixed investments
	fmt.Println("\n=== Fixed Investments ===")
	investments, err := s.SmartLending.GetFixedInvestments("", "USDT", 1, 10, 0)
	if err != nil {
		fmt.Println("get investments error:", err)
	} else {
		fmt.Printf("Active investments: %d\n", len(investments.Records))
		for _, inv := range investments.Records {
			fmt.Printf("  ID: %s, Amount: %s, Interest Paid: %s\n",
				inv.ID, inv.Amount, inv.InterestPaid)
		}
	}

	// Get interest payment history
	fmt.Println("\n=== Interest Payment History ===")
	interestHistory, err := s.SmartLending.GetFixedInterestHistory("", "USDT", 10, 0)
	if err != nil {
		fmt.Println("interest history error:", err)
	} else {
		fmt.Printf("Interest payments: %d\n", len(interestHistory.Records))
		for _, payment := range interestHistory.Records {
			fmt.Printf("  Plan: %s, Investment: %s, Amount: %s %s\n",
				payment.PlanID, payment.InvestmentID, payment.Amount, payment.Ticker)
		}
	}

	// ============ Flexible Lending ============
	fmt.Println("\n========== Flexible Lending ==========")

	// Get available flexible lending plans
	fmt.Println("\n=== Flexible Lending Plans ===")
	flexPlans, err := s.SmartLending.GetFlexPlans("USDT", 10, 0)
	if err != nil {
		fmt.Println("flex plans error:", err)
	} else {
		fmt.Printf("Available flex plans: %d\n", len(flexPlans))
		for _, plan := range flexPlans {
			fmt.Printf("  ID: %s, Ticker: %s, Max Rate: %s%%\n",
				plan.ID, plan.Ticker, plan.MaxRate)
			fmt.Printf("    Min: %s, Max: %s\n", plan.MinInvestment, plan.MaxInvestment)
		}
	}

	// Create flexible investment
	fmt.Println("\n=== Create Flexible Investment ===")
	if len(flexPlans) > 0 {
		flexInv, err := s.SmartLending.FlexInvest(flexPlans[0].ID, "50")
		if err != nil {
			fmt.Println("flex invest error:", err)
		} else {
			fmt.Printf("Flex Investment ID: %s, Invested: %s, Auto-reinvest: %v\n",
				flexInv.ID, flexInv.Invested, flexInv.WithAutoInvest)
		}
	}

	// Get flexible investments
	fmt.Println("\n=== Flexible Investments ===")
	flexInvestments, err := s.SmartLending.GetFlexInvestments("USDT", "", "", 1, 10, 0)
	if err != nil {
		fmt.Println("get flex investments error:", err)
	} else {
		fmt.Printf("Active flex investments: %d\n", len(flexInvestments.Data))
		for _, inv := range flexInvestments.Data {
			fmt.Printf("  ID: %s, Currency: %s, Invested: %s, Auto-reinvest: %v\n",
				inv.ID, inv.Currency, inv.Invested, inv.WithAutoInvest)
		}
	}

	// Enable auto-reinvest
	fmt.Println("\n=== Toggle Auto-Reinvest ===")
	if len(flexInvestments.Data) > 0 {
		updated, err := s.SmartLending.FlexSetAutoInvest(flexInvestments.Data[0].ID, true)
		if err != nil {
			fmt.Println("set auto-invest error:", err)
		} else {
			fmt.Printf("Auto-reinvest enabled: %v\n", updated.WithAutoInvest)
		}
	}

	// Get flexible investment history
	fmt.Println("\n=== Flexible Investment History ===")
	flexHistory, err := s.SmartLending.GetFlexHistory("", "", "", 0, 0, nil, 10, 0)
	if err != nil {
		fmt.Println("flex history error:", err)
	} else {
		fmt.Printf("History records: %d\n", len(flexHistory.Data))
		for _, record := range flexHistory.Data {
			fmt.Printf("  Investment: %s, Amount: %s %s, Action: %d\n",
				record.InvestmentID, record.Amount, record.Currency, record.ActionType)
		}
	}
}
