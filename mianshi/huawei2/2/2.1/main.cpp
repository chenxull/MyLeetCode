#include<iostream>
#include<string>
#include<vector>

using namespace std;

const int MAXN = 1010;

string all[MAXN];
int cnt = 0;

int notblank(char c)
{
	if (c >= '0'&&c <= '9') return 1;
	else if (c >= 'a'&&c <= 'z') return 1;
	else if (c >= 'A'&&c <= 'Z') return 1;
	else if (c == '-') return 1;
	else return 0;

}

int main()
{
	string str;
	while (cin>>str)
	{
		while (str[0] == '-') { str.erase(0, 1); }
		while (str[str.length() - 1] == '-') { str.erase(str.length() - 1, 1); }
		int p = 0;
		int num = 0;
		for (int i = 0; i < str.length(); i++)
		{
			if (!notblank(str[i]))
			{
				string temp = str.substr(p, num);
				all[cnt++] = temp;
				num = 0;
				p = i + 1;
			}
			if (str[i] == '-' && str[i + 1] == '-')
			{
				string temp = str.substr(p, num);
				all[cnt++] = temp;
				num = 0;
				p = i + 2;
			}
			num++;
		}
		if (p <= str.length() - 1)
		{
			string temp = str.substr(p, str.length()-p);
			all[cnt++] = temp;
		}
	}
	for (int i = cnt - 1; i >= 0; i--)
	{
		cout << all[i];
		if (i != 0) cout << " ";
	}
	//system("pause");
	return 0;
}