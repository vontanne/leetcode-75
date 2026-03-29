package queue

/*
933. Number of Recent Calls

You have a RecentCounter class which counts the number of recent requests within a certain time frame.

Implement the RecentCounter class:
- RecentCounter() Initializes the counter with zero recent requests.
- int ping(int t) Adds a new request at time t, where t represents some time in milliseconds,
  and returns the number of requests that has happened in the past 3000 milliseconds
  (including the new request). Specifically, return the number of requests that have
  happened in the inclusive range [t - 3000, t].

It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.

Example 1:
	Input
	["RecentCounter", "ping", "ping", "ping", "ping"]
	[[], [1], [100], [3001], [3002]]
	Output
	[null, 1, 2, 3, 3]

	Explanation
	RecentCounter recentCounter = new RecentCounter();
	recentCounter.ping(1);     // requests = [1], range is [-2999,1], return 1
	recentCounter.ping(100);   // requests = [1, 100], range is [-2900,100], return 2
	recentCounter.ping(3001);  // requests = [1, 100, 3001], range is [1,3001], return 3
	recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002], range is [2,3002], return 3

Constraints:
	- 1 <= t <= 10^9
	- Each test case will call ping with strictly increasing values of t.
	- At most 10^4 calls will be made to ping.
*/

type RecentCounter struct {
	timestamps []int
	head       int
}

func Constructor() RecentCounter {
	return RecentCounter{
		timestamps: make([]int, 0, 10000),
		head:       0,
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.timestamps = append(rc.timestamps, t)

	for rc.head < len(rc.timestamps) && rc.timestamps[rc.head] < t-3000 {
		rc.head++
	}

	return len(rc.timestamps) - rc.head
}

/*
Time Complexity: O(1) amortized per ping
	- Each request is added once and head moves forward at most once per element
	- Total operations across all pings is O(n) where n is number of pings

Space Complexity: O(n)
	- Pre-allocated slice stores all timestamps (up to 10^4)
	- No memory released during execution, but bounded by constraint
*/
