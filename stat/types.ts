export interface Stat {
  Bhattacharyya: (xs: number[], ys: number[]) => number;
  BivariateMoment: (
    q: number,
    p: number,
    xs: number[],
    ys: number[],
    weights: number[],
  ) => number;
  CDF: (q: number, xs: number[], weights: number[]) => number;
  ChiSquare: (xs: number[], ys: number[]) => number;
  CircularMean: (xs: number[], weights: number[]) => number;
  Correlation: (xs: number[], ys: number[], weights: number[]) => number;
  Covariance: (xs: number[], ys: number[], weights: number[]) => number;
  CrossEntropy: (xs: number[], ys: number[]) => number;
  Entropy: (xs: number[]) => number;
  ExKurtosis: (xs: number[], weights: number[]) => number;
  GeometricMean: (xs: number[], weights: number[]) => number;
  HarmonicMean: (xs: number[], weights: number[]) => number;
  Hellinger: (xs: number[], ys: number[]) => number;
  Histogram: (
    counts: number[],
    divs: number[],
    xs: number[],
    bins: number,
  ) => number[];
  JensenShannon: (xs: number[], ys: number[]) => number;
  Kendall: (xs: number[], ys: number[], weights: number[]) => number;
  KolmogorovSmirnov: (
    xs: number[],
    xw: number[],
    ys: number[],
    yw: number[],
  ) => number;
  KullbackLeibler: (xs: number[], ys: number[]) => number;
  LinearRegression: (
    xs: number[],
    ys: number[],
    weights: number[],
    origin: boolean,
  ) => { alpha: number; beta: number };
  Mean: (xs: number[], weights: number[]) => number;
  MeanStdDev: (
    xs: number[],
    weights: number[],
  ) => { mean: number; stdDev: number };
  MeanVariance: (
    xs: number[],
    weights: number[],
  ) => { mean: number; variance: number };
  Mode: (xs: number[], weights: number[]) => { value: number; count: number };
  Moment: (q: number, xs: number[], weights: number[]) => number;
  MomentAbout: (
    q: number,
    xs: number[],
    mean: number,
    weights: number[],
  ) => number;
  PopMeanStdDev: (
    xs: number[],
    weights: number[],
  ) => { mean: number; stdDev: number };
  PopMeanVariance: (
    xs: number[],
    weights: number[],
  ) => { mean: number; variance: number };
  PopStdDev: (xs: number[], weights: number[]) => number;
  PopVariance: (xs: number[], weights: number[]) => number;
  Quantile: (
    q: number,
    type: "linear" | "empirical",
    xs: number[],
    weights: number[],
  ) => number;
  RNoughtSquared: (
    xs: number[],
    ys: number[],
    weights: number[],
    beta: number,
  ) => number;
  ROC: (
    cutoffs: number[],
    ys: number[],
    classes: boolean[],
    weights: number[],
  ) => { tpr: number[]; fpr: number[]; tresh: number[] };
  RSquared: (
    xs: number[],
    ys: number[],
    weights: number[],
    alpha: number,
    beta: number,
  ) => number;
  RSquaredFrom: (
    estimates: number[],
    ys: number[],
    weights: number[],
  ) => number;
  Skew: (xs: number[], weights: number[]) => number;
  SortWeighted: (xs: number[], weights: number[]) => number[];
  SortWeightedLabeled: (
    xs: number[],
    labels: boolean[],
    weights: number[],
  ) => { xs: number[]; labels: boolean[] };
  StdDev: (xs: number[], weights: number[]) => number;
  StdErr: (std: number, size: number) => number;
  StdScore: (x: number, mean: number, stdDev: number) => number;
  TOC: (
    classes: boolean[],
    ys: number[],
  ) => { min: number[]; ntp: number[]; max: number[] };
  Variance: (xs: number[], weights: number[]) => number;
}
