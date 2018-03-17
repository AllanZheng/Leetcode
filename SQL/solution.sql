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
select class from courses group by class having count(Distinct student)>=5 