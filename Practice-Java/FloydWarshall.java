import java.io.*;
import java.util.*;

public class FloydWarshall {

    private static final int INF = Integer.MAX_VALUE;

    public static int[][] readGraph(String filePath) throws IOException {
        try (BufferedReader br = new BufferedReader(new FileReader(filePath))) {
            // Read number of nodes
            int numNodes = Integer.parseInt(br.readLine().trim());
            int[][] adjMatrix = new int[numNodes][numNodes];

            // Initialize adjacency matrix with infinity
            for (int i = 0; i < numNodes; i++) {
                Arrays.fill(adjMatrix[i], INF);
                adjMatrix[i][i] = 0; // Distance from node to itself is zero
            }

            // Read edges and their costs
            String line;
            while ((line = br.readLine()) != null) {
                String[] parts = line.trim().split(" ");
                int src = Integer.parseInt(parts[0]);
                int dest = Integer.parseInt(parts[1]);
                int cost = Integer.parseInt(parts[2]);
                adjMatrix[src][dest] = cost;
                adjMatrix[dest][src] = cost; // For undirected graph
            }

            return adjMatrix;
        }
    }

    public static int[][] floydWarshall(int[][] adjMatrix) {
        int numNodes = adjMatrix.length;
        int[][] distMatrix = new int[numNodes][numNodes];

        // Initialize the distance matrix with the adjacency matrix
        for (int i = 0; i < numNodes; i++) {
            System.arraycopy(adjMatrix[i], 0, distMatrix[i], 0, numNodes);
        }

        // Apply Floyd-Warshall algorithm
        for (int k = 0; k < numNodes; k++) {
            for (int i = 0; i < numNodes; i++) {
                for (int j = 0; j < numNodes; j++) {
                    if (distMatrix[i][k] != INF && distMatrix[k][j] != INF) {
                        distMatrix[i][j] = Math.min(distMatrix[i][j], distMatrix[i][k] + distMatrix[k][j]);
                    }
                }
            }
        }

        return distMatrix;
    }

    public static void writeCostMatrix(String filePath, int[][] distMatrix) throws IOException {
        BufferedWriter bw = new BufferedWriter(new FileWriter(filePath));

        for (int[] row : distMatrix) {
            for (int value : row) {
                if (value == INF) {
                    bw.write("INF ");
                } else {
                    bw.write(value + " ");
                }
            }
            bw.newLine();
        }

        bw.close();
    }

    public static void main(String[] args) {
        String inputFile = "C:\\Users\\VM-Eduardo\\Desktop\\Huertas\\Asignacion03\\SolaJava\\graph.txt"; // Input file path
        String outputFile = "C:\\Users\\VM-Eduardo\\Desktop\\Huertas\\Asignacion03\\SolaJava\\cost_matrix.txt"; // Output file path

        try {
            // Step 1: Read the graph from the file
            int[][] adjMatrix = readGraph(inputFile);

            // Step 2: Apply the Floyd-Warshall algorithm
            int[][] costMatrix = floydWarshall(adjMatrix);

            // Step 3: Write the cost matrix to the output file
            writeCostMatrix(outputFile, costMatrix);

            System.out.println("Cost matrix has been written to " + outputFile);
        } catch (IOException e) {
            System.err.println("An error occurred: " + e.getMessage());
        }
    }
}