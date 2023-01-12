# Base
### status
- Health check
- /status
    - `GET`


# API

### Login
- Login (Token 발행)
- /api/chokomemo/login
    - `POST`
    - Request
        - __loginId__: `string`
        - __password__: `string`
        ```json
        {
            "loginId": "testId",
            "password": "password"
        }
        ```
    - Response
        - __error__: `any`, if error occured
        - __userId__: `string`
        - __token__: `string`
        ```json
        {
            "userId": "1234",
            "token": "temp_token"
        }

        // error
        {
            "error": { "message": "..." }
        }
        ```

### Get All Memo List
- 모든 메모 리스트 가져오기
- api/chokomemo/all-memos
    - `GET`
    - Request
        - __userId__: `string`,
        - __token__: `string`
        ```json
        {
            "userId": "1234",
            "token": "temp_token"
        }
        ```
    - Response
        - __error__: `any`, if error occured
        - __memoList__ `array` of `memoListInfo`
            - __memoListInfo__: `object`
                - __memoId__: `number`
                - __title__: `string`
        ```json
        {
            "memoList": [
                {
                    "memoId": 1,
                    "title": "memo-title",
                }
            ]
        }

        // error
        {
            "error": { "message": "..." }
        }
        ```

### Get Memo
- 특정 메모 가져오기
- api/chokomemo/memo
    - `GET`
    - Request
        - __userId__: `string`,
        - __token__: `string`,
        - __memoId__: `number`
        ```json
        {
            "userId": "1234",
            "token": "temp_token",
            "memoId": 1
        }
        ```
    - Response
        - __error__: `any`, if error occured
        - __title__: `string`
        - __content__: `string`
        - __createdDate__: `datetime`
        - __lastUpdatedDate__: `datetime`
        ```json
        {
            "title": "memo-title",
            "content": "blah blah",
            "createdDate": "",
            "lastUpdatedDate": ""
        }

        // error
        {
            "error": { "message": "..." }
        }
        ```

### Create Memo
- 새로운 메모 생성
- /api/chokomemo/memo
    - `POST`
    - Request
        - __userId__: `string`
        - __token__: `string`
        - __title__: `string`
        - __content__: `string`
        ```json
        {
            "userId": "1234",
            "token": "temp_token",
            "title": "memo-title",
            "content": "blah blah"
        }
        ```
    - Response
        - __error__: `any`, if error occured
        - __isSuccess__: `boolean`
        - __memoId__: `number`
        ```json
        {
            "isSuccess": true,
            "memoId": 1,
        }

        // error
        {
            "error": { "message": "..." }
        }
        ```

### Update Memo
- 기존 메모 수정
- /api/chokomemo/memo
    - `PUT`
    - Request
        - __userId__: `string`
        - __token__: `string`
        - __memoId__: `number`
        - __title__: `string`
        - __content__: `string`
        ```json
        {
            "userId": "1234",
            "token": "temp_token",
            "memoId": 1,
            "title": "memo-title",
            "content": "blah blah"
        }
        ```
    - Response
        - __error__: `any`, if error occured
        - __isSuccess__: `boolean`
        ```json
        {
            "isSuccess": true
        }

        // error
        {
            "error": { "message": "..." }
        }
        ```

### Delete Memo
- 기존 메모 삭제
- /api/chokomemo/memo
    - `DELETE`
    - Request
        - __userId__: `string`
        - __token__: `string`
        - __memoIds__: `array` of `number`
        ```json
        {
            "userId": "1234",
            "token": "temp_token",
            "memoIds": [1, 2, 3]
        }
        ```
    - Response
        - __error__: `any`, if error occured
        - __deletedMemoIds__: `array` of `number`
        ```json
        {
            "deletedMemoIds": [1, 2, 3]
        }

        // error
        {
            "error": { "message": "..." }
        }
        ```