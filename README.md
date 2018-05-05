# CSV-normalization
Go was my language of choice because of how it handles UTF-8. I wrote some code in Python and was getting unicodes (e.g. U+1F34F) everywhere.

## Normalization
The entire CSV is in the UTF-8 character set and the following fields were normalized:

- [x] The Timestamp column should be formatted in ISO-8601 format.
- [x] The Timestamp column should be assumed to be in US/Pacific time; please convert it to US/Eastern.
- [x] All ZIP codes should be formatted as 5 digits. If there are less than 5 digits, assume 0 as the prefix.
- [x] All name columns should be converted to uppercase. There will be non-English names.
- [x] The Address column should be passed through as is, except for Unicode validation. Please note there are commas in the Address field; your CSV parsing will need to take that into account. Commas will only be present inside a quoted string.
- [x] The columns `FooDuration` and `BarDuration` are in HH:MM:SS.MS format (where MS is milliseconds); please convert them to a floating point seconds format.
- [ ] The column "TotalDuration" is filled with garbage data. For each row, please replace the value of TotalDuration with the sum of FooDuration and BarDuration.
- [ ] The column "Notes" is free form text input by end-users; please do not perform any transformations on this column. If there are invalid UTF-8 characters, please replace them with the Unicode Replacement Character.

## Running the program
Read in *sample.csv* via stdin:
```
go run main.go < sample.csv
```

Read in *sample-with-broken-utf8.csv* via stdin:
```
go run main.go < sample-with-broken-utf8.csv
```

## Notes
* The following lines were removed from each of the files for testing:
  * sample.csv - 11/11/11 11:11:11 AM,überTown,10001,Prompt Negotiator,1:23:32.123,1:32:33.123,zzsasdfa,"I’m just gonna say, this is AMAZING. WHAT NEGOTIATIONS."
  * sample-with-broken-utf8.cs - 11/11/11 11:11:11 AM,√ºberTown,10001,Prompt Negotiator,1:23:32.123,1:32:33.123,zzsasdfa,"I‚Äôm just gonna say, this is AMAZING. WHAT NEGOTIATIONS."
