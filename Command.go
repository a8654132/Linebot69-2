package main

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"log"
)

var RedisPort = "6379"
var RedisIP = "140.115.153.185"
var mac = "8c:a9:82:03:3d:68"
func Redis_Get(KEY_NAME string) ([]byte, error){
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
    CheckError(err)
    defer c.Close()

	v, err2 := redis.Bytes(c.Do("GET", KEY_NAME))
	return v, err2
}

func Redis_Set(KEY_NAME string, VAL []byte) error{
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
    CheckError(err)
    defer c.Close()

	_, err2 := c.Do("SET", KEY_NAME, VAL)
	return err2
}

func Redis_Del(KEY_NAME string) error{
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
	CheckError(err)
	defer c.Close()

	_, err2 := c.Do("DEL", KEY_NAME)
	CheckError(err2)
	return err2
}


func Redis_DelAllUser()error{
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
	CheckError(err)
	defer c.Close()

	user := Redis_AllUser()
	var err2 error = nil
	for i:=0 ; i<len(user) ; i++{
		_, err2 = c.Do("DEL", user[i])
		CheckError(err2)
		if err2 != nil{
			return err2
		}
	}
	return err2
}


//this will find all user with a strings slice
func Redis_AllUser()([]string){
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
	CheckError(err)
	defer c.Close()

	user, err2 := redis.Strings(c.Do("KEYS", "*"))
	CheckError(err2)
	return user
}

//if not find, return is fail
func Redis_FindUser(USER string)(bool){
	RedisIPPORT := fmt.Sprintf("%s:%s", RedisIP, RedisPort)
	c, err := redis.Dial("tcp", RedisIPPORT)
	CheckError(err)
	defer c.Close()

	EXIST, err2 := redis.Bool(c.Do("EXISTS", USER))
	CheckError(err2)
	return EXIST
}



func CheckError(err error) {
    if err  != nil {
        log.Println("Error: " , err)
        // os.Exit(0)
    }
}
