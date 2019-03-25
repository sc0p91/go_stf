package main

type entity struct {
	alive	bool
	player	bool
	name    string
	class   string
	hp      int
	mp      int
	maxHp   int
	maxMp   int
	exp     int
	maxExp  int
	lvl     int
	sta     int
	str     int
	int     int
	dex     int
	attacks [4]string
	items   [2]string
}

func (p *entity) selectClass(name string, sta int, str int, dex int, int int) {
	p.class = name
	p.sta = sta
	p.str = str
	p.dex = dex
	p.int = int
}