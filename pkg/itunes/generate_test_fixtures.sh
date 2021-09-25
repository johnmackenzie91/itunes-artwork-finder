curl -i https://itunes.apple.com/search?country=gb&entity=album&term=the+killers+-+sams+town \
  --output ./testdata/sams_town

curl -i https://itunes.apple.com/search?country=gb\u0026entity=album\u0026term=the+killers+-+sams+town \
  --output ./testdata/bad_request

curl -i https://itunes.apple.com/search?country=gb&entity=album&term=ddwed+-+fwefwefwefwf \
  --output ./testdata/no_results
