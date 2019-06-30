package problem950

import (
	"fmt"
	"sort"
)

/*
 * @lc aspp=leetcode id=950 lang=golang
 *
 * [950] Reveal Cards In Increasing Order
 *
 * https://leetcode.com/problems/reveal-cards-in-increasing-order/description/
 *
 * algorithms
 * Medium (71.17%)
 * Total Accepted:    12.3K
 * Total Submissions: 17.2K
 * Testcase Example:  '[17,13,11,2,3,5,7]'
 *
 * In a deck of cards, every card has a unique integer.  You can order the deck
 * in any order you want.
 *
 * Initially, all the cards start face down (unrevealed) in one deck.
 *
 * Now, you do the following steps repeatedly, until all cards are
 * revealed:
 *
 *
 * Take the top card of the deck, reveal it, and take it out of the deck.
 * If there are still cards in the deck, put the next top card of the deck at
 * the bottom of the deck.
 * If there are still unrevealed cards, go back to step 1.  Otherwise, stop.
 *
 *
 * Return an ordering of the deck that would reveal the cards in increasing
 * order.
 *
 * The first entry in the answer is considered to be the top of the deck.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [17,13,11,2,3,5,7]
 * Output: [2,13,3,11,5,17,7]
 * Explanation:
 * We get the deck in the order [17,13,11,2,3,5,7] (this order doesn't matter),
 * and reorder it.
 * After reordering, the deck starts as [2,13,3,11,5,17,7], where 2 is the top
 * of the deck.
 * We reveal 2, and move 13 to the bottom.  The deck is now [3,11,5,17,7,13].
 * We reveal 3, and move 11 to the bottom.  The deck is now [5,17,7,13,11].
 * We reveal 5, and move 17 to the bottom.  The deck is now [7,13,11,17].
 * We reveal 7, and move 13 to the bottom.  The deck is now [11,17,13].
 * We reveal 11, and move 17 to the bottom.  The deck is now [13,17].
 * We reveal 13, and move 17 to the bottom.  The deck is now [17].
 * We reveal 17.
 * Since all the cards revealed are in increasing order, the answer is
 * correct.
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 1000
 * 1 <= A[i] <= 10^6
 * A[i] != A[j] for all i != j
 *
 *
 *
 *
 */
func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	size := len(deck)

	t := make([]int, 0, size*2)
	b := 0

	for i := size - 1; i >= 0; i-- {
		t = append(t, deck[i])
		if i > 0 {
			t = append(t, t[b])
			b++
			fmt.Println("DEBUG1::t", t)
		}

	}

	for i, j := b, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
		fmt.Println("DEBUG2::t", t)
	}

	return t[b:]
}
