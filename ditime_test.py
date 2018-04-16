#!/usr/bin/env python
#coding:utf-8
import time
import unittest
from ditime import *

class Test_ditime(unittest.TestCase):
	def test_ToShortTime(self):
		self.assertEqual(ToShortTime("0808"), ("0808",None))
		self.assertEqual(ToShortTime("2017-08-08"), ("0808",None))
		self.assertEqual(ToShortTime("2017-08-08T23:54:23+09:00"), ("0808",None))
		self.assertEqual(ToShortTime("a"), ("a","약속한 시간포멧 형태가 아닙니다."))
	def test_ToNormalTime(self):
		self.assertEqual(ToNormalTime("0808"), ("%s-08-08" %  (time.strftime("%Y")),None))
		self.assertEqual(ToNormalTime("2017-08-08"), ("2017-08-08",None))
		self.assertEqual(ToNormalTime("2017-08-08T23:54:23+09:00"), ("2017-08-08",None))
		self.assertEqual(ToNormalTime("a"), ("a","약속한 시간포멧 형태가 아닙니다."))
	def test_ToFullTime(self):
		self.assertEqual(ToFullTime("0808"), ("%s-08-08T19:00:00+09:00" %  (time.strftime("%Y")),None))
		self.assertEqual(ToFullTime("2017-08-08"), ("2017-08-08T19:00:00+09:00",None))
		self.assertEqual(ToFullTime("2017-08-08T23:54:23+09:00"), ("2017-08-08T23:54:23+09:00",None))
		self.assertEqual(ToFullTime("a"), ("a","약속한 시간포멧 형태가 아닙니다."))
	def test_ToExcelTime(self):
		self.assertEqual(ToExcelTime("0808"), ("08/08",None))
		self.assertEqual(ToExcelTime("2017-08-08"), ("08/08",None))
		self.assertEqual(ToExcelTime("2017-08-08T23:54:23+09:00"), ("08/08",None))
		self.assertEqual(ToExcelTime("a"), ("a","약속한 시간포멧 형태가 아닙니다."))

if __name__ == "__main__":
	unittest.main()
