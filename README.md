# Question
a team organizing a university career fair has a list of companies along with their respective arrival time and duration to stay.  
only one company can present at any time. given each company arrival time and duration they will stay.  
determine the maximum number of presentation that can be hosted during the career fair

constraint  
n = number of companies 1 <= n <= 50  
1 < arrival[i] < 1000    
1 < duration[i] < 1000    
both arrival and duration must have same element


# Example Input & Output
n = 5  
arrival = [1, 3, 3, 5, 7]  
duration = [2, 2, 1, 2, 1]  

output: 4  
the first company arrives at time 1 and stay for 2 hours. at time 3 two companies arrive. but only 1 can stay for either 1 or 2 hours. the next companies arrive at times 5 and 7 and do not conflict with any others. in total there can be maximum  of 4 presentation
