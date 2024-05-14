# toml2sql
toml2sql은 altibase input plugin for telegraf에 포함되어 있는 query.toml 파일을 query.sql 파일로 변환해주는 유틸리티이다.   
## 개요
telegraf에서 알티베이스로 수행하는 sql이 문제가 없는지 검증이 필요하다.   
이 sql은 query.toml 파일에 정의되어 있다.   
따라서 query.toml 파일에서 sql을 추출하여 테스트해보면 결과를 검증할 수 있다.   
toml2sql을 수행하면 현재 디렉토리의 query.toml 파일을 query.sql로 변환한다.
## 사용법
```bash
./toml2sql
is -f query.sql -o query.out
grep -i err query.out
```
## 참고사항
query.toml 파일은 최신 버전의 altibase input plugin에 포함된 파일로 테스트 하는 것을 권장한다.    
참고로 본 리파지토리의 query.toml은 altibase input plugin v1.0.0 패키지에 포함된 파일이다.
