# Install the 'boot' package if not already installed
# install.packages("boot")

library(boot)

# Define a statistic to be estimated
statistic <- function(data, indices) {
  # Median function to apply on resampled data
  return(median(data[indices]))
}

# Example data
data <- c(8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68)

# Apply bootstrap sampling to estimate the standard error of the median
result <- boot(data, statistic, R = 1000)

# Display the result
print(result)
