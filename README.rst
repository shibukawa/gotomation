Gotomation
============

Automation library like ``java.awt.Robot``.

Document
----------

https://godoc.org/github.com/shibukawa/gotomation

License
--------

Apache-2

Known Issue
-----------

* Mouse coursor moving is not working on VirtualBox host OS with mouse cursor integration.
* Windows: mouse cursor moving is not working (I tested on Windows 10 Pro on VirtualBox)
* X11: String typing is not implemented. In future, I will add this only for cgo environment.
  XCB (and xgb - golang version) doesn't support ``XStringToKeysym()`` equivalent function.

Thanks
------

* https://github.com/BurntSushi/xgb

  It uses ``xgb`` for X11 environment.

* https://github.com/go-vgo/robotgo

  This code is inspired by robotogo. I rewrite into pure go, add utf-8 support etc as much as possible.
