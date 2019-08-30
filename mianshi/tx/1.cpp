#include <cstdio>
#include <algorithm>
#include <iostream>
using namespace std;


int a[10010],b[10010];
int n,m;
int main(){
    cin>>n;
    for (int i=0;i<n;i++){
        cin>>a[i];
    }
    cin>>m;
    for (int i=0;i<m;i++){
        cin>>b[i];
    }

    sort(a,a+n);
    sort(b,b+m);

    int ans=0;
    int posa=0,posb=0;

    while (posa<n && posb<m){
        if (a[posa]<=b[posb]){
        ans++;
        posa++;
        posb++;
        }else {
            posb++;
        }
    }
    cout<<ans<<endl;
}