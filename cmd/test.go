package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

import "github.com/go-redis/redis"
import "math/rand"
import "time"
import "strconv"

var ip string
var port string
var status string
var concurrent int
var verbose bool
var cluster bool
var total int

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	seed := strconv.FormatInt((time.Now().Unix()), 10)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	result := string(seed) + string(b)
	return result
}

func NewClusterClient(ip string, k string, v string) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{ip + ":" + port},
		Password: "",
	})
	for i := 0; i < total; i++ {
		k = randSeq(6)
		if verbose == true {
                        fmt.Println("set key:", k, "| value:", v)
                }
		err := client.Set(k, v, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func NewClient(ip string, k string, v string) {
	client := redis.NewClient(&redis.Options{
		Addr:     ip + ":" + port,
		Password: "",
		DB:       0,
	})

	for i := 1; i < total; i++ {
		k = randSeq(6)
		if verbose == true {
			fmt.Println("set key:", k, "| value:", v)
		}
		err := client.Set(k, v, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func MultiThreadBench(n int) {
	n = concurrent
	fmt.Println(time.Now())
	chs := make([]chan int, n)
	for i := 0; i < n; i++ {
		chs[i] = make(chan int)
		go func(ch chan int, i int) {
			if cluster == true {
			NewClusterClient(ip, randSeq(6), "xxx")
			} else if cluster == false {
			NewClient(ip, randSeq(6), "xxx")
			}
			ch <- 1
		}(chs[i], i)
	}
	for _, ch := range chs {
		<-ch
	}
	fmt.Println(time.Now())
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("test called")
		if status == "" {
			cmd.Help()
			return
		} else if status == "print" {
			fmt.Print("WIP\n")
			return
		}
		fmt.Println("[INFO] start: ", ip, port)
		//fmt.Println(verbose)
		MultiThreadBench(concurrent)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//testCmd.Flags().BoolP("verbose", "v", true, "verbose mode")
	testCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	testCmd.PersistentFlags().BoolVarP(&cluster, "cluster", "C", false, "cluster mode on/off")
	testCmd.Flags().StringVarP(&ip, "ip", "i", "127.0.0.1", "ip of redis server")
	testCmd.Flags().StringVarP(&port, "port", "p", "6379", "port of redis server")
	testCmd.Flags().StringVarP(&status, "status", "s", "", "[start|print] start test now! (required)")
	testCmd.Flags().IntVarP(&concurrent, "concurrent", "c", 500, "number of concurrent clients")
	testCmd.Flags().IntVarP(&total, "request", "n", 200, "number of requests")
}
