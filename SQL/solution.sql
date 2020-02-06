#627. Swap Salary
Update salary set sex=IF(sex='m','f','m')
# 620. Not Boring Movies
select id,movie,description,rating from cinema where id%2!=0 AND description!='boring' order by rating desc
# Dulicate emails
select Email
from Person
group by Email
having count(*)>1;
#175 combine_two_tables
select person.FirstName, person.LastName, address.City, address.State
from Person left join Address on person.personId=address.personId
#181
select a.Name as Employee from employee as a inner join employee as copy on a.ManagerId=copy.Id AND A.salary>copy.salary
# 183
select Name as Customers from Customers where NOT EXISTS (select * from Orders where Customers.Id=Orders.CustomerId)
#596. Classes More Than 5 distinct Students 
select class from courses group by class having count(Distinct student)>=5 、

#20200207 1179. Reformat Department Table (Unsolved)
select id,
	sum(case when month = 'jan' then revenue else null end) as Jan_Revenue,
	sum(case when month = 'feb' then revenue else null end) as Feb_Revenue,
	sum(case when month = 'mar' then revenue else null end) as Mar_Revenue,
	sum(case when month = 'apr' then revenue else null end) as Apr_Revenue,
	sum(case when month = 'may' then revenue else null end) as May_Revenue,
	sum(case when month = 'jun' then revenue else null end) as Jun_Revenue,
	sum(case when month = 'jul' then revenue else null end) as Jul_Revenue,
	sum(case when month = 'aug' then revenue else null end) as Aug_Revenue,
	sum(case when month = 'sep' then revenue else null end) as Sep_Revenue,
	sum(case when month = 'oct' then revenue else null end) as Oct_Revenue,
	sum(case when month = 'nov' then revenue else null end) as Nov_Revenue,
	sum(case when month = 'dec' then revenue else null end) as Dec_Revenue
from department
group by id
order by id