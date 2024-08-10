import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
  } from 'chart.js';
import { Bar } from 'react-chartjs-2';

ChartJS.register(
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
  );

interface HistogramSortChartProps {
    dataInput: number[];
  }

  const HistogramSortChart: React.FC<HistogramSortChartProps> = ({ dataInput }) => {

    const data = {
        labels: dataInput,
        datasets: [
          {
            data: dataInput,
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            borderColor: 'rgba(75, 192, 192, 1)',
            borderWidth: 1,
          },
        ],
      };

      const options = {
        responsive: true,
        maintainAspectRatio: false,
        borderRadius: 10,
        scales: {
          y: {
            grid: {
                offset: false
            }
          },
        },
      };


    return (
        <Bar data={data} options={options} />
    )
  };


export { HistogramSortChart }
