library(tidyverse)
library(here)

# part 1

data_file <- read_lines(here::here("Data", "05.txt"))  %>%
  as.data.frame() %>%
  rename(full = 1) %>%
  separate(full, sep = " -> ", into = c("start", "end")) %>%
  mutate(idx = 1:n()) %>%
  pivot_longer(c(-idx)) %>%
  separate(value, sep = ",", into = c("x", "y")) %>%
  mutate(across(c(x, y), as.numeric)) %>%
  group_by(idx) %>%
  mutate(min_max_x = min(x) - max(x), min_max_y = min(y) - max(y)) %>%
  ungroup() %>%
  filter(min_max_x == 0 | min_max_y == 0) %>%
  pivot_longer(c(x, y), names_to = "coord", values_to = "coord_val") %>%
  mutate(name = glue::glue("{name}_{coord}")) %>%
  select(-coord, -min_max_x, -min_max_y) %>%
  pivot_wider(names_from = "name", values_from = "coord_val") %>%
  rowwise() %>%
  mutate(seq_x = list(start_x:end_x), seq_y = list(start_y:end_y)) %>%
  ungroup() %>%
  select(idx, seq_x, seq_y) %>%
  unnest(c(seq_x, seq_y))

data_file %>%
  count(seq_x, seq_y) %>%
  filter(n > 1) %>%
  count()

# part 2

data_file_raw <- read_lines(here::here("Data", "05.txt"))  %>%
  as.data.frame() %>%
  rename(full = 1) %>%
  separate(full, sep = " -> ", into = c("start", "end")) %>%
  mutate(idx = 1:n()) %>%
  pivot_longer(c(-idx)) %>%
  separate(value, sep = ",", into = c("x", "y")) %>%
  mutate(across(c(x, y), as.numeric)) %>%
  pivot_longer(c(x, y), names_to = "coord", values_to = "coord_val") %>%
  mutate(name = glue::glue("{name}_{coord}")) %>%
  select(-coord) %>%
  pivot_wider(names_from = "name", values_from = "coord_val")

data_file_raw %>%
  rowwise() %>%
  mutate(seq_x = list(start_x:end_x), seq_y = list(start_y:end_y)) %>%
  select(idx, seq_x, seq_y) %>%
  unnest(c(seq_x, seq_y)) %>%
  count(seq_x, seq_y) %>%
  filter(n > 1) %>%
  count()
