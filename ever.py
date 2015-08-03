#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from datetime import datetime

with open("/tmp/cron_times_test","ab+") as f:
    time_now = datetime.now().strftime("%H:%M:%S")
    f.write("write {0}\n".format(time_now).encode('utf-8'))
