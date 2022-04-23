package leetc

type Twitter struct {
	us      map[int]*user
	curTime int
}

type user struct {
	tweets   []*tweet
	followee map[int]*user
}

type tweet struct {
	tweetId  int
	postTime int
}

func Constructor355() Twitter {
	return Twitter{us: make(map[int]*user)}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.addUserIfAbsent(userId)
	this.us[userId].tweets = append(this.us[userId].tweets, &tweet{tweetId: tweetId, postTime: this.getCurTime()})
}

func (this *Twitter) GetNewsFeed(userId int) (res []int) {
	this.addUserIfAbsent(userId)
	users := append([]*user{}, this.us[userId])
	for _, v := range this.us[userId].followee {
		users = append(users, v)
	}
	idx := make([]int, len(users))
	for i, u := range users {
		idx[i] = len(u.tweets) - 1
	}
	for i := 10; i > 0; i-- {
		lrt, ui := 0, -1
		for j, u := range users {
			if idx[j] < 0 {
				continue
			}
			tweet := u.tweets[idx[j]]
			if tweet.postTime > lrt {
				ui = j
				lrt = tweet.postTime
			}
		}
		if ui == -1 {
			return
		}
		res = append(res, users[ui].tweets[idx[ui]].tweetId)
		idx[ui]--
	}
	return
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.addUserIfAbsent(followerId)
	this.addUserIfAbsent(followeeId)
	if _, exist := this.us[followeeId].followee[followeeId]; !exist {
		this.us[followerId].followee[followeeId] = this.us[followeeId]
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, exist := this.us[followerId]; !exist {
		return
	}
	delete(this.us[followerId].followee, followeeId)
}

func (this *Twitter) addUserIfAbsent(userId int) {
	if _, exist := this.us[userId]; !exist {
		this.us[userId] = &user{followee: make(map[int]*user)}
	}
}

func (this *Twitter) getCurTime() int {
	this.curTime++
	return this.curTime
}
