# synwork-processor-generic

This project provides a [synwork.io](http://www.synwork.io) processor for handling json structure.

configure the processor

```
synwork {
    required_processor  {
        source = "sbl.systems/sappi/generic"
        version = ">= 0.0.1, < 1.0.0"
    }

 }
```

Find out more details create a minimal configuration file generic.snw with this required processor,
go in this directory and
>>> synwork init
>>> synwork help -p generic

