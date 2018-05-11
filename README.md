<table><tr><td><b>Name</b></td><td><b>Zerolog</b></td><td><b>ZapSugar</b></td><td><b>Zap</b></td><td><b>Onelog</b></td><td><b>Logrus</b></td></tr>
<tr><td colspan=6><b>Enabled/WithContext/Fields</b></td></tr>
<tr><td>Floats32-8</td><td>&#x1F538; <b>245ns ±13%</b></td><td>430ns ± 7%</td><td>372ns ± 3%</td><td>n/a</td><td>3.16µs ±13%</td></tr>
<tr><td>Object-8</td><td>80.5ns ± 2%</td><td>302ns ± 3%</td><td>270ns ± 6%</td><td>&#x1F538; <b>79.2ns ± 2%</b></td><td>2.68µs ± 4%</td></tr>
<tr><td>Strings-8</td><td>&#x1F538; <b>169ns ± 1%</b></td><td>789ns ± 2%</td><td>712ns ± 4%</td><td>n/a</td><td>3.63µs ± 4%</td></tr>
<tr><td>Ints-8</td><td>&#x1F538; <b>82.0ns ± 4%</b></td><td>307ns ± 4%</td><td>241ns ± 3%</td><td>n/a</td><td>2.91µs ± 3%</td></tr>
<tr><td>Int-8</td><td>59.5ns ± 6%</td><td>255ns ± 5%</td><td>206ns ± 3%</td><td>&#x1F538; <b>53.7ns ± 1%</b></td><td>2.81µs ± 4%</td></tr>
<tr><td>Float32-8</td><td>&#x1F538; <b>75.5ns ± 1%</b></td><td>273ns ± 4%</td><td>226ns ± 1%</td><td>n/a</td><td>2.72µs ± 7%</td></tr>
<tr><td>String-8</td><td>&#x1F538; <b>69.0ns ± 7%</b></td><td>332ns ± 1%</td><td>278ns ± 3%</td><td>84.5ns ± 2%</td><td>2.76µs ± 4%</td></tr>
<tr><td>Bools-8</td><td>&#x1F538; <b>63.4ns ± 4%</b></td><td>288ns ± 8%</td><td>218ns ± 4%</td><td>n/a</td><td>3.02µs ± 1%</td></tr>
<tr><td>Error-8</td><td>&#x1F538; <b>63.3ns ± 3%</b></td><td>294ns ± 2%</td><td>248ns ± 2%</td><td>73.6ns ± 8%</td><td>n/a</td></tr>
<tr><td>Time-8</td><td>&#x1F538; <b>129ns ±22%</b></td><td>449ns ± 4%</td><td>388ns ± 4%</td><td>n/a</td><td>n/a</td></tr>
<tr><td>Errors-8</td><td>&#x1F538; <b>135ns ±16%</b></td><td>n/a</td><td>n/a</td><td>n/a</td><td>n/a</td></tr>
<tr><td>Float64-8</td><td>94.6ns ± 4%</td><td>298ns ± 3%</td><td>236ns ± 4%</td><td>&#x1F538; <b>78.8ns ± 2%</b></td><td>2.88µs ± 8%</td></tr>
<tr><td>Floats64-8</td><td>&#x1F538; <b>383ns ±12%</b></td><td>552ns ± 2%</td><td>479ns ± 5%</td><td>n/a</td><td>3.05µs ± 4%</td></tr>
<tr><td>Bool-8</td><td>56.9ns ± 4%</td><td>244ns ± 1%</td><td>208ns ± 5%</td><td>&#x1F538; <b>52.8ns ±10%</b></td><td>2.75µs ± 5%</td></tr>
<tr><td colspan=6><b>Enabled/WithContext</b></td></tr>
<tr><td>Msg-8</td><td>45.7ns ± 3%</td><td>229ns ± 2%</td><td>143ns ± 8%</td><td>&#x1F538; <b>43.5ns ± 2%</b></td><td>2.10µs ± 6%</td></tr>
<tr><td>MsgComplex-8</td><td>71.3ns ± 7%</td><td>275ns ± 5%</td><td>177ns ±10%</td><td>&#x1F538; <b>61.4ns ± 3%</b></td><td>2.27µs ± 7%</td></tr>
<tr><td>Formatting-8</td><td>&#x1F538; <b>266ns ±15%</b></td><td>327ns ± 4%</td><td>n/a</td><td>n/a</td><td>2.38µs ± 7%</td></tr>
<tr><td colspan=6><b>Enabled/NoContext/Fields</b></td></tr>
<tr><td>Bool-8</td><td>&#x1F538; <b>49.5ns ± 3%</b></td><td>243ns ± 3%</td><td>199ns ± 2%</td><td>49.7ns ± 3%</td><td>1.59µs ± 2%</td></tr>
<tr><td>Int-8</td><td>53.3ns ± 2%</td><td>254ns ± 6%</td><td>203ns ± 3%</td><td>&#x1F538; <b>50.3ns ± 2%</b></td><td>1.68µs ± 6%</td></tr>
<tr><td>Float64-8</td><td>86.2ns ± 2%</td><td>280ns ± 3%</td><td>233ns ± 2%</td><td>&#x1F538; <b>76.3ns ± 3%</b></td><td>1.63µs ± 7%</td></tr>
<tr><td>Ints-8</td><td>&#x1F538; <b>76.9ns ± 3%</b></td><td>303ns ± 2%</td><td>239ns ± 3%</td><td>n/a</td><td>1.74µs ± 3%</td></tr>
<tr><td>Strings-8</td><td>&#x1F538; <b>158ns ± 3%</b></td><td>780ns ± 4%</td><td>700ns ± 6%</td><td>n/a</td><td>2.52µs ±10%</td></tr>
<tr><td>Time-8</td><td>&#x1F538; <b>102ns ± 2%</b></td><td>446ns ± 4%</td><td>441ns ± 6%</td><td>n/a</td><td>n/a</td></tr>
<tr><td>String-8</td><td>&#x1F538; <b>60.4ns ±20%</b></td><td>326ns ± 3%</td><td>304ns ± 4%</td><td>82.2ns ± 2%</td><td>1.71µs ± 4%</td></tr>
<tr><td>Errors-8</td><td>&#x1F538; <b>123ns ± 2%</b></td><td>n/a</td><td>n/a</td><td>n/a</td><td>n/a</td></tr>
<tr><td>Error-8</td><td>&#x1F538; <b>61.9ns ± 5%</b></td><td>297ns ± 2%</td><td>271ns ± 2%</td><td>66.1ns ± 2%</td><td>n/a</td></tr>
<tr><td>Floats32-8</td><td>&#x1F538; <b>216ns ± 3%</b></td><td>423ns ± 4%</td><td>361ns ± 2%</td><td>n/a</td><td>1.90µs ± 5%</td></tr>
<tr><td>Float32-8</td><td>&#x1F538; <b>61.7ns ± 1%</b></td><td>269ns ± 3%</td><td>241ns ± 2%</td><td>n/a</td><td>1.63µs ± 3%</td></tr>
<tr><td>Bools-8</td><td>&#x1F538; <b>57.8ns ± 2%</b></td><td>282ns ± 4%</td><td>222ns ±11%</td><td>n/a</td><td>1.87µs ± 3%</td></tr>
<tr><td>Object-8</td><td>75.9ns ± 6%</td><td>302ns ± 3%</td><td>263ns ± 4%</td><td>&#x1F538; <b>75.3ns ± 3%</b></td><td>1.64µs ± 3%</td></tr>
<tr><td>Floats64-8</td><td>&#x1F538; <b>346ns ± 2%</b></td><td>548ns ± 2%</td><td>476ns ± 2%</td><td>n/a</td><td>2.06µs ± 3%</td></tr>
<tr><td colspan=6><b>Enabled/NoContext</b></td></tr>
<tr><td>Msg-8</td><td>&#x1F538; <b>37.5ns ± 3%</b></td><td>228ns ± 1%</td><td>151ns ± 9%</td><td>39.9ns ± 1%</td><td>1.14µs ± 5%</td></tr>
<tr><td>MsgComplex-8</td><td>57.0ns ± 4%</td><td>272ns ± 3%</td><td>182ns ± 5%</td><td>&#x1F538; <b>56.7ns ± 0%</b></td><td>1.13µs ± 5%</td></tr>
<tr><td>Formatting-8</td><td>&#x1F538; <b>178ns ± 4%</b></td><td>318ns ± 4%</td><td>n/a</td><td>n/a</td><td>1.33µs ± 7%</td></tr>
<tr><td colspan=6><b>Disabled/WithContext/Fields</b></td></tr>
<tr><td>Object-8</td><td>6.24ns ± 2%</td><td>14.6ns ±18%</td><td>29.1ns ± 3%</td><td>&#x1F538; <b>4.36ns ± 6%</b></td><td>276ns ± 3%</td></tr>
<tr><td>Ints-8</td><td>&#x1F538; <b>6.53ns ± 2%</b></td><td>27.5ns ± 4%</td><td>44.8ns ± 1%</td><td>n/a</td><td>297ns ± 3%</td></tr>
<tr><td>Errors-8</td><td>&#x1F538; <b>6.49ns ± 5%</b></td><td>n/a</td><td>n/a</td><td>n/a</td><td>n/a</td></tr>
<tr><td>Bool-8</td><td>6.12ns ± 4%</td><td>14.7ns ± 7%</td><td>29.4ns ± 5%</td><td>&#x1F538; <b>2.02ns ± 1%</b></td><td>276ns ± 2%</td></tr>
<tr><td>Bools-8</td><td>&#x1F538; <b>6.59ns ± 6%</b></td><td>29.3ns ± 5%</td><td>45.1ns ± 2%</td><td>n/a</td><td>290ns ± 2%</td></tr>
<tr><td>Floats64-8</td><td>&#x1F538; <b>6.37ns ± 4%</b></td><td>28.5ns ± 2%</td><td>44.2ns ± 2%</td><td>n/a</td><td>288ns ± 2%</td></tr>
<tr><td>Float32-8</td><td>&#x1F538; <b>5.96ns ± 3%</b></td><td>20.0ns ± 4%</td><td>28.9ns ± 1%</td><td>n/a</td><td>283ns ± 1%</td></tr>
<tr><td>String-8</td><td>7.20ns ± 1%</td><td>24.5ns ± 5%</td><td>30.5ns ± 6%</td><td>&#x1F538; <b>2.97ns ± 3%</b></td><td>287ns ± 4%</td></tr>
<tr><td>Floats32-8</td><td>&#x1F538; <b>6.69ns ±14%</b></td><td>28.2ns ± 5%</td><td>44.7ns ± 2%</td><td>n/a</td><td>300ns ± 3%</td></tr>
<tr><td>Strings-8</td><td>&#x1F538; <b>6.55ns ± 3%</b></td><td>28.6ns ± 1%</td><td>44.7ns ± 2%</td><td>n/a</td><td>289ns ± 2%</td></tr>
<tr><td>Float64-8</td><td>8.68ns ± 2%</td><td>22.6ns ± 4%</td><td>31.4ns ± 2%</td><td>&#x1F538; <b>2.44ns ± 4%</b></td><td>313ns ±19%</td></tr>
<tr><td>Time-8</td><td>&#x1F538; <b>7.04ns ±21%</b></td><td>28.7ns ± 3%</td><td>44.8ns ± 1%</td><td>n/a</td><td>n/a</td></tr>
<tr><td>Error-8</td><td>6.25ns ± 2%</td><td>15.1ns ± 9%</td><td>30.2ns ± 7%</td><td>&#x1F538; <b>2.96ns ± 3%</b></td><td>n/a</td></tr>
<tr><td>Int-8</td><td>6.18ns ± 4%</td><td>20.5ns ± 5%</td><td>31.7ns ± 8%</td><td>&#x1F538; <b>2.24ns ± 8%</b></td><td>284ns ± 3%</td></tr>
<tr><td colspan=6><b>Disabled/WithContext</b></td></tr>
<tr><td>Msg-8</td><td>4.58ns ± 3%</td><td>13.9ns ± 8%</td><td>7.34ns ± 4%</td><td>&#x1F538; <b>1.07ns ± 2%</b></td><td>21.7ns ± 6%</td></tr>
<tr><td>MsgComplex-8</td><td>4.48ns ± 4%</td><td>13.7ns ± 7%</td><td>7.29ns ± 3%</td><td>&#x1F538; <b>1.07ns ± 1%</b></td><td>21.7ns ± 5%</td></tr>
<tr><td>Formatting-8</td><td>4.90ns ± 1%</td><td>3.43ns ± 2%</td><td>n/a</td><td>n/a</td><td>&#x1F538; <b>2.41ns ± 3%</b></td></tr>
<tr><td colspan=6><b>Disabled/NoContext/Fields</b></td></tr>
<tr><td>Time-8</td><td>&#x1F538; <b>6.75ns ± 3%</b></td><td>29.5ns ± 5%</td><td>44.3ns ± 2%</td><td>n/a</td><td>n/a</td></tr>
<tr><td>Object-8</td><td>6.03ns ± 3%</td><td>15.3ns ±14%</td><td>29.4ns ± 2%</td><td>&#x1F538; <b>4.68ns ± 7%</b></td><td>250ns ± 2%</td></tr>
<tr><td>Int-8</td><td>6.38ns ± 3%</td><td>26.1ns ±41%</td><td>30.6ns ± 4%</td><td>&#x1F538; <b>1.99ns ± 2%</b></td><td>255ns ± 3%</td></tr>
<tr><td>String-8</td><td>7.22ns ± 3%</td><td>27.2ns ±23%</td><td>31.1ns ± 5%</td><td>&#x1F538; <b>2.32ns ± 5%</b></td><td>271ns ± 3%</td></tr>
<tr><td>Bools-8</td><td>&#x1F538; <b>6.84ns ± 8%</b></td><td>30.6ns ±12%</td><td>44.3ns ± 2%</td><td>n/a</td><td>285ns ±19%</td></tr>
<tr><td>Float64-8</td><td>8.77ns ± 7%</td><td>23.9ns ± 8%</td><td>30.6ns ± 2%</td><td>&#x1F538; <b>2.00ns ± 1%</b></td><td>260ns ± 2%</td></tr>
<tr><td>Floats64-8</td><td>&#x1F538; <b>6.40ns ± 2%</b></td><td>28.5ns ± 7%</td><td>44.5ns ± 2%</td><td>n/a</td><td>271ns ±14%</td></tr>
<tr><td>Errors-8</td><td>&#x1F538; <b>6.33ns ± 2%</b></td><td>n/a</td><td>n/a</td><td>n/a</td><td>n/a</td></tr>
<tr><td>Bool-8</td><td>6.19ns ± 2%</td><td>15.2ns ±13%</td><td>30.1ns ± 8%</td><td>&#x1F538; <b>1.97ns ± 3%</b></td><td>251ns ± 4%</td></tr>
<tr><td>Float32-8</td><td>&#x1F538; <b>6.14ns ± 4%</b></td><td>19.9ns ± 3%</td><td>29.5ns ± 6%</td><td>n/a</td><td>256ns ± 4%</td></tr>
<tr><td>Ints-8</td><td>&#x1F538; <b>6.46ns ± 2%</b></td><td>29.1ns ± 3%</td><td>45.2ns ± 3%</td><td>n/a</td><td>267ns ± 2%</td></tr>
<tr><td>Error-8</td><td>6.32ns ± 5%</td><td>15.9ns ±20%</td><td>29.6ns ± 4%</td><td>&#x1F538; <b>2.32ns ± 4%</b></td><td>n/a</td></tr>
<tr><td>Floats32-8</td><td>&#x1F538; <b>6.50ns ± 4%</b></td><td>29.1ns ± 7%</td><td>45.2ns ± 3%</td><td>n/a</td><td>263ns ± 2%</td></tr>
<tr><td>Strings-8</td><td>&#x1F538; <b>6.88ns ± 7%</b></td><td>29.6ns ± 6%</td><td>46.0ns ± 5%</td><td>n/a</td><td>268ns ± 3%</td></tr>
<tr><td colspan=6><b>Disabled/NoContext</b></td></tr>
<tr><td>Msg-8</td><td>4.40ns ± 3%</td><td>18.0ns ±20%</td><td>7.26ns ± 2%</td><td>&#x1F538; <b>1.08ns ± 0%</b></td><td>25.8ns ± 5%</td></tr>
<tr><td>MsgComplex-8</td><td>4.49ns ± 2%</td><td>13.9ns ± 9%</td><td>7.27ns ± 2%</td><td>&#x1F538; <b>1.08ns ± 3%</b></td><td>25.8ns ± 4%</td></tr>
<tr><td>Formatting-8</td><td>5.03ns ± 3%</td><td>4.19ns ±25%</td><td>n/a</td><td>n/a</td><td>&#x1F538; <b>2.59ns ± 9%</b></td></tr>
</table>
