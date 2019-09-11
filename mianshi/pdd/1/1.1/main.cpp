#include<iostream>
#include<algorithm>
using namespace std;

const int MAXN = int(1e2+10);

int n;
int a[MAXN];

bool cmp(int a, int b)
{
	if (a % 2 == 0 && b % 2 != 0) return true;
	else if (a % 2 == 0 && b % 2 == 0)
	{
		if (a > b) return true;
		else return false;
	}
	else if (a % 2 != 0 && b % 2 == 0) return false;
	else
	{
		if (a > b) return true;
		else return false;
	}
}


int main()
{
	int x;
	int cnt = 0;
	char f;
	while (cin >> x)
	{
		a[cnt++] = x;
		cin >> f;
		if (f == ';')
		{
			cin >> n;
			break;
		}
	}
	sort(a, a + n, cmp);
	for (int i = 0; i < n; i++)
	{
		cout << a[i];
		if (i != n - 1) cout << ",";
	}
	//system("pause");
	return 0;
}