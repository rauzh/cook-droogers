import pandas as pd
import matplotlib.pyplot as plt

# Читаем данные из CSV-файла
data = pd.read_csv('data_restore')

# Извлекаем данные
x = data['X']
y1 = data['Y1']
y2 = data['Y2']
y3 = data['Y3']
y4 = data['Y4']

# Создаем гистограммы
plt.figure(figsize=(10, 6))

plt.grid(color='grey', linestyle='--', linewidth=1, alpha=0.2)

# Наслаивание столбцов с прозрачностью
plt.bar(x, y4, color='purple', width=0.5, label='1500 тыс. записей', alpha=0.4)
plt.bar(x, y3, color='brown', width=0.5, label='1000 тыс. записей', alpha=0.6)
plt.bar(x, y2, color='lightgrey', width=0.5, label='500 тыс. записей', alpha=0.8)
plt.bar(x, y1, color='black', width=0.5, label='200 тыс. записей', alpha=1)

# Настройки графика
plt.xlabel('Число потоков, шт.')
plt.ylabel('Время выполнения, мс')
plt.title('Зависимость времени выполнения восстановления данных от числа потоков')
plt.legend()
plt.xticks(list(range(len(x)+1)))
plt.tight_layout()

# Показываем график
plt.show()