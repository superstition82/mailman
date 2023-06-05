#!/bin/bash

num_addresses=10000

domains=("google.com" "naver.com" "hanmail.net" "kakao.com")
output_file="random_emails.txt"
num_domains=${#domains[@]}

for ((i=1; i<=num_addresses; i++))
do
  domain_index=$(( RANDOM % num_domains ))

  username=$(LC_CTYPE=C tr -dc 'a-z0-9' < /dev/urandom | fold -w 3 | head -n 1)
  domain=${domains[domain_index]}

  email="$username@$domain"

  echo $email >> $output_file
done

echo "Random email addresses generated and saved in $output_file."
