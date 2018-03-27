#!/usr/bin/env python
# coding:utf-8

import re
import time

def ToShortTime(t):
	"""
	2017-10-25T16:21:47+09:00 또는 2017-10-25 형태의 날짜를 1025로 바꾸어준다.
	"""
	if re.search('\d{4}$',t):
		return t, None
	if re.search('\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$', t):
		return t[5:10].replace("-",""), None
	if re.search('\d{4}-\d{2}-\d{2}', t):
		return t[5:7]+t[8:10], None
	return t, "약속한 시간포멧 형태가 아닙니다."

def ToNormalTime(t):
	"""
	1025 또는 2017-10-25T16:21:47+09:00 형태의 날짜를 2017-10-25로 바꾸어준다.
	"""
	if re.search('\d{4}-\d{2}-\d{2}$', t):
		return t, None
	if re.search('\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$', t):
		return t[:10], None
	if re.search('\d{4}$',t):
		return '%s-%s-%s' % (time.strftime("%Y"), t[0:2], t[2:4]), None
	return t, "약속한 시간포멧 형태가 아닙니다."

def ToFullTime(t):
	"""
	1025 또는 2017-10-25 형태의 날짜를 2017-10-25T19:00:00+09:00 형태로 바꾸어준다.
	"""
	if re.search('\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$', t):
		return t, None
	if re.search('\d{4}$',t):
		# 1025처럼 짧은 문자열은 한국 퇴근시간으로 바꾼다.
		return '%s-%s-%sT19:00:00+09:00' % (time.strftime("%Y"), t[0:2], t[2:4]), None
	if re.search('\d{4}-\d{2}-\d{2}$', t):
		# 2017-10-25 문자열은 한국 퇴근시간으로 바꾼다.
		return '%sT19:00:00+09:00' % (t), None
	return t, "약속한 시간포멧 형태가 아닙니다."

if __name__ == "__main__":
	print ToNormalTime("2017-08-08T23:45:23+09:00")
