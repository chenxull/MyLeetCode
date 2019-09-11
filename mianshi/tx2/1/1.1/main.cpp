#include<bits/stdc++.h>
using namespace std;
int main()
 
{
    int n,m,odd_n=0,even_n=0;
 
    int c,k,odd_m=0,even_m=0;
    cin>>n>>m;
    for(int i=0;i<n;i++){
        cin>>c;
        if(c%2)odd_n++;
        else even_n++;
    }
 
    for(int i=0;i<m;i++){
        cin>>k;
        if(k%2)odd_m++;
        else even_m++;
    }
 
    cout<<min(odd_n,even_m)+min(odd_m,even_n);
}