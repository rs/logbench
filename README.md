| Name | Zerolog | ZapSugar | Zap | Onelog | Logrus |
|-|-|-|-|-|-|
| Enabled/WithContext/MsgComplex-8 | 71.3ns ± 7% | 275ns ± 5% | 177ns ±10% | &#x1F538; **61.4ns ± 3%** | 2.27µs ± 7% |
| Enabled/WithContext/Msg-8 | 45.7ns ± 3% | 229ns ± 2% | 143ns ± 8% | &#x1F538; **43.5ns ± 2%** | 2.10µs ± 6% |
| Enabled/WithContext/Formatting-8 | &#x1F538; **266ns ±15%** | 327ns ± 4% | n/a | n/a | 2.38µs ± 7% |
| Enabled/WithContext/Fields/Time-8 | &#x1F538; **129ns ±22%** | 449ns ± 4% | 388ns ± 4% | n/a | n/a |
| Enabled/WithContext/Fields/Strings-8 | &#x1F538; **169ns ± 1%** | 789ns ± 2% | 712ns ± 4% | n/a | 3.63µs ± 4% |
| Enabled/WithContext/Fields/String-8 | &#x1F538; **69.0ns ± 7%** | 332ns ± 1% | 278ns ± 3% | 84.5ns ± 2% | 2.76µs ± 4% |
| Enabled/WithContext/Fields/Object-8 | 80.5ns ± 2% | 302ns ± 3% | 270ns ± 6% | &#x1F538; **79.2ns ± 2%** | 2.68µs ± 4% |
| Enabled/WithContext/Fields/Ints-8 | &#x1F538; **82.0ns ± 4%** | 307ns ± 4% | 241ns ± 3% | n/a | 2.91µs ± 3% |
| Enabled/WithContext/Fields/Int-8 | 59.5ns ± 6% | 255ns ± 5% | 206ns ± 3% | &#x1F538; **53.7ns ± 1%** | 2.81µs ± 4% |
| Enabled/WithContext/Fields/Floats64-8 | &#x1F538; **383ns ±12%** | 552ns ± 2% | 479ns ± 5% | n/a | 3.05µs ± 4% |
| Enabled/WithContext/Fields/Floats32-8 | &#x1F538; **245ns ±13%** | 430ns ± 7% | 372ns ± 3% | n/a | 3.16µs ±13% |
| Enabled/WithContext/Fields/Float64-8 | 94.6ns ± 4% | 298ns ± 3% | 236ns ± 4% | &#x1F538; **78.8ns ± 2%** | 2.88µs ± 8% |
| Enabled/WithContext/Fields/Float32-8 | &#x1F538; **75.5ns ± 1%** | 273ns ± 4% | 226ns ± 1% | n/a | 2.72µs ± 7% |
| Enabled/WithContext/Fields/Errors-8 | &#x1F538; **135ns ±16%** | n/a | n/a | n/a | n/a |
| Enabled/WithContext/Fields/Error-8 | &#x1F538; **63.3ns ± 3%** | 294ns ± 2% | 248ns ± 2% | 73.6ns ± 8% | n/a |
| Enabled/WithContext/Fields/Bools-8 | &#x1F538; **63.4ns ± 4%** | 288ns ± 8% | 218ns ± 4% | n/a | 3.02µs ± 1% |
| Enabled/WithContext/Fields/Bool-8 | 56.9ns ± 4% | 244ns ± 1% | 208ns ± 5% | &#x1F538; **52.8ns ±10%** | 2.75µs ± 5% |
| Enabled/NoContext/MsgComplex-8 | 57.0ns ± 4% | 272ns ± 3% | 182ns ± 5% | &#x1F538; **56.7ns ± 0%** | 1.13µs ± 5% |
| Enabled/NoContext/Msg-8 | &#x1F538; **37.5ns ± 3%** | 228ns ± 1% | 151ns ± 9% | 39.9ns ± 1% | 1.14µs ± 5% |
| Enabled/NoContext/Formatting-8 | &#x1F538; **178ns ± 4%** | 318ns ± 4% | n/a | n/a | 1.33µs ± 7% |
| Enabled/NoContext/Fields/Time-8 | &#x1F538; **102ns ± 2%** | 446ns ± 4% | 441ns ± 6% | n/a | n/a |
| Enabled/NoContext/Fields/Strings-8 | &#x1F538; **158ns ± 3%** | 780ns ± 4% | 700ns ± 6% | n/a | 2.52µs ±10% |
| Enabled/NoContext/Fields/String-8 | &#x1F538; **60.4ns ±20%** | 326ns ± 3% | 304ns ± 4% | 82.2ns ± 2% | 1.71µs ± 4% |
| Enabled/NoContext/Fields/Object-8 | 75.9ns ± 6% | 302ns ± 3% | 263ns ± 4% | &#x1F538; **75.3ns ± 3%** | 1.64µs ± 3% |
| Enabled/NoContext/Fields/Ints-8 | &#x1F538; **76.9ns ± 3%** | 303ns ± 2% | 239ns ± 3% | n/a | 1.74µs ± 3% |
| Enabled/NoContext/Fields/Int-8 | 53.3ns ± 2% | 254ns ± 6% | 203ns ± 3% | &#x1F538; **50.3ns ± 2%** | 1.68µs ± 6% |
| Enabled/NoContext/Fields/Floats64-8 | &#x1F538; **346ns ± 2%** | 548ns ± 2% | 476ns ± 2% | n/a | 2.06µs ± 3% |
| Enabled/NoContext/Fields/Floats32-8 | &#x1F538; **216ns ± 3%** | 423ns ± 4% | 361ns ± 2% | n/a | 1.90µs ± 5% |
| Enabled/NoContext/Fields/Float64-8 | 86.2ns ± 2% | 280ns ± 3% | 233ns ± 2% | &#x1F538; **76.3ns ± 3%** | 1.63µs ± 7% |
| Enabled/NoContext/Fields/Float32-8 | &#x1F538; **61.7ns ± 1%** | 269ns ± 3% | 241ns ± 2% | n/a | 1.63µs ± 3% |
| Enabled/NoContext/Fields/Errors-8 | &#x1F538; **123ns ± 2%** | n/a | n/a | n/a | n/a |
| Enabled/NoContext/Fields/Error-8 | &#x1F538; **61.9ns ± 5%** | 297ns ± 2% | 271ns ± 2% | 66.1ns ± 2% | n/a |
| Enabled/NoContext/Fields/Bools-8 | &#x1F538; **57.8ns ± 2%** | 282ns ± 4% | 222ns ±11% | n/a | 1.87µs ± 3% |
| Enabled/NoContext/Fields/Bool-8 | &#x1F538; **49.5ns ± 3%** | 243ns ± 3% | 199ns ± 2% | 49.7ns ± 3% | 1.59µs ± 2% |
| Disabled/WithContext/MsgComplex-8 | 4.48ns ± 4% | 13.7ns ± 7% | 7.29ns ± 3% | &#x1F538; **1.07ns ± 1%** | 21.7ns ± 5% |
| Disabled/WithContext/Msg-8 | 4.58ns ± 3% | 13.9ns ± 8% | 7.34ns ± 4% | &#x1F538; **1.07ns ± 2%** | 21.7ns ± 6% |
| Disabled/WithContext/Formatting-8 | 4.90ns ± 1% | 3.43ns ± 2% | n/a | n/a | &#x1F538; **2.41ns ± 3%** |
| Disabled/WithContext/Fields/Time-8 | &#x1F538; **7.04ns ±21%** | 28.7ns ± 3% | 44.8ns ± 1% | n/a | n/a |
| Disabled/WithContext/Fields/Strings-8 | &#x1F538; **6.55ns ± 3%** | 28.6ns ± 1% | 44.7ns ± 2% | n/a | 289ns ± 2% |
| Disabled/WithContext/Fields/String-8 | 7.20ns ± 1% | 24.5ns ± 5% | 30.5ns ± 6% | &#x1F538; **2.97ns ± 3%** | 287ns ± 4% |
| Disabled/WithContext/Fields/Object-8 | 6.24ns ± 2% | 14.6ns ±18% | 29.1ns ± 3% | &#x1F538; **4.36ns ± 6%** | 276ns ± 3% |
| Disabled/WithContext/Fields/Ints-8 | &#x1F538; **6.53ns ± 2%** | 27.5ns ± 4% | 44.8ns ± 1% | n/a | 297ns ± 3% |
| Disabled/WithContext/Fields/Int-8 | 6.18ns ± 4% | 20.5ns ± 5% | 31.7ns ± 8% | &#x1F538; **2.24ns ± 8%** | 284ns ± 3% |
| Disabled/WithContext/Fields/Floats64-8 | &#x1F538; **6.37ns ± 4%** | 28.5ns ± 2% | 44.2ns ± 2% | n/a | 288ns ± 2% |
| Disabled/WithContext/Fields/Floats32-8 | &#x1F538; **6.69ns ±14%** | 28.2ns ± 5% | 44.7ns ± 2% | n/a | 300ns ± 3% |
| Disabled/WithContext/Fields/Float64-8 | 8.68ns ± 2% | 22.6ns ± 4% | 31.4ns ± 2% | &#x1F538; **2.44ns ± 4%** | 313ns ±19% |
| Disabled/WithContext/Fields/Float32-8 | &#x1F538; **5.96ns ± 3%** | 20.0ns ± 4% | 28.9ns ± 1% | n/a | 283ns ± 1% |
| Disabled/WithContext/Fields/Errors-8 | &#x1F538; **6.49ns ± 5%** | n/a | n/a | n/a | n/a |
| Disabled/WithContext/Fields/Error-8 | 6.25ns ± 2% | 15.1ns ± 9% | 30.2ns ± 7% | &#x1F538; **2.96ns ± 3%** | n/a |
| Disabled/WithContext/Fields/Bools-8 | &#x1F538; **6.59ns ± 6%** | 29.3ns ± 5% | 45.1ns ± 2% | n/a | 290ns ± 2% |
| Disabled/WithContext/Fields/Bool-8 | 6.12ns ± 4% | 14.7ns ± 7% | 29.4ns ± 5% | &#x1F538; **2.02ns ± 1%** | 276ns ± 2% |
| Disabled/NoContext/MsgComplex-8 | 4.49ns ± 2% | 13.9ns ± 9% | 7.27ns ± 2% | &#x1F538; **1.08ns ± 3%** | 25.8ns ± 4% |
| Disabled/NoContext/Msg-8 | 4.40ns ± 3% | 18.0ns ±20% | 7.26ns ± 2% | &#x1F538; **1.08ns ± 0%** | 25.8ns ± 5% |
| Disabled/NoContext/Formatting-8 | 5.03ns ± 3% | 4.19ns ±25% | n/a | n/a | &#x1F538; **2.59ns ± 9%** |
| Disabled/NoContext/Fields/Time-8 | &#x1F538; **6.75ns ± 3%** | 29.5ns ± 5% | 44.3ns ± 2% | n/a | n/a |
| Disabled/NoContext/Fields/Strings-8 | &#x1F538; **6.88ns ± 7%** | 29.6ns ± 6% | 46.0ns ± 5% | n/a | 268ns ± 3% |
| Disabled/NoContext/Fields/String-8 | 7.22ns ± 3% | 27.2ns ±23% | 31.1ns ± 5% | &#x1F538; **2.32ns ± 5%** | 271ns ± 3% |
| Disabled/NoContext/Fields/Object-8 | 6.03ns ± 3% | 15.3ns ±14% | 29.4ns ± 2% | &#x1F538; **4.68ns ± 7%** | 250ns ± 2% |
| Disabled/NoContext/Fields/Ints-8 | &#x1F538; **6.46ns ± 2%** | 29.1ns ± 3% | 45.2ns ± 3% | n/a | 267ns ± 2% |
| Disabled/NoContext/Fields/Int-8 | 6.38ns ± 3% | 26.1ns ±41% | 30.6ns ± 4% | &#x1F538; **1.99ns ± 2%** | 255ns ± 3% |
| Disabled/NoContext/Fields/Floats64-8 | &#x1F538; **6.40ns ± 2%** | 28.5ns ± 7% | 44.5ns ± 2% | n/a | 271ns ±14% |
| Disabled/NoContext/Fields/Floats32-8 | &#x1F538; **6.50ns ± 4%** | 29.1ns ± 7% | 45.2ns ± 3% | n/a | 263ns ± 2% |
| Disabled/NoContext/Fields/Float64-8 | 8.77ns ± 7% | 23.9ns ± 8% | 30.6ns ± 2% | &#x1F538; **2.00ns ± 1%** | 260ns ± 2% |
| Disabled/NoContext/Fields/Float32-8 | &#x1F538; **6.14ns ± 4%** | 19.9ns ± 3% | 29.5ns ± 6% | n/a | 256ns ± 4% |
| Disabled/NoContext/Fields/Errors-8 | &#x1F538; **6.33ns ± 2%** | n/a | n/a | n/a | n/a |
| Disabled/NoContext/Fields/Error-8 | 6.32ns ± 5% | 15.9ns ±20% | 29.6ns ± 4% | &#x1F538; **2.32ns ± 4%** | n/a |
| Disabled/NoContext/Fields/Bools-8 | &#x1F538; **6.84ns ± 8%** | 30.6ns ±12% | 44.3ns ± 2% | n/a | 285ns ±19% |
| Disabled/NoContext/Fields/Bool-8 | 6.19ns ± 2% | 15.2ns ±13% | 30.1ns ± 8% | &#x1F538; **1.97ns ± 3%** | 251ns ± 4% |
